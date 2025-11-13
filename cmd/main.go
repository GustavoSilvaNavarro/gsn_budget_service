package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gsn_budget_service/api/server"
	"github.com/gsn_budget_service/internal"
	"github.com/gsn_budget_service/internal/config"
	"github.com/gsn_budget_service/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()

	// Global ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize application with all dependencies (DB, queries, etc.)
	appConns, err := internal.New(ctx, config.Cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("☠️  Failed to initialize application...")
	}
	defer appConns.CloseAppConnections()

	srv := server.StartServer(appConns)
	srvErr := make(chan error, 1)

	go func() {
		log.Info().Msgf(
			"🚀 %s API service is running, listening on PORT: %d", strings.ToUpper(config.Cfg.NAME), config.Cfg.PORT,
		)

		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			srvErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-srvErr:
		log.Fatal().Err(err).Msg("💥 Server crashed")
	case sig := <-quit:
		log.Warn().Msgf("🧹 Caught signal: %s — shutting down gracefully...", sig.String())
		cancel()

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("⚠️ Server forced to shutdown")
			os.Exit(1)
		}

		log.Info().Msg("✅ Server stopped gracefully")
	}
}
