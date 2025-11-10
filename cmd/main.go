package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gsn_budget_service/api/server"
	"github.com/gsn_budget_service/internal/config"
	"github.com/gsn_budget_service/internal/db"
	"github.com/gsn_budget_service/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()

	// Global ctx
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	srv := server.StartServer(config.Cfg)
	srvErr := make(chan error, 1)

	go func() {
		dbConnection, err := db.StartDbConnection(ctx, config.Cfg)
		if err != nil {
			srvErr <- fmt.Errorf("☠️  DB connection failed... | Error: %w", err)
			return
		}
		defer dbConnection.Close()

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
