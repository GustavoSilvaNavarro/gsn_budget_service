package logger

import (
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gsn_budget_service/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	buildInfo, _ := debug.ReadBuildInfo()
	zerolog.TimeFieldFormat = time.RFC3339

	logLevel := strings.ToLower(config.Cfg.LOG_LEVEL)
	envLevel := strings.ToLower(config.Cfg.ENVIRONMENT)

	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.DebugLevel // in case of error we default to debug
	}

	zerolog.SetGlobalLevel(level)

	log.Logger = log.Logger.With().
		Caller(). // Caller (file and line number) to every log event
		Str("env", envLevel).
		Str("service", config.Cfg.NAME).
		Int("pid", os.Getpid()).
		Str("lang", "Golang").
		Str("go_version", buildInfo.GoVersion).
		Logger()

	log.Info().Msg("🪵 Logger initialized!")
}
