package logger

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/gsn_budget_service/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var once sync.Once

func InitLogger() {
	once.Do(func() {
		buildInfo, _ := debug.ReadBuildInfo()
		zerolog.TimeFieldFormat = time.RFC3339

		logLevel := strings.ToLower(config.Cfg.LOG_LEVEL)
		envLevel := strings.ToLower(config.Cfg.ENVIRONMENT)

		level, err := zerolog.ParseLevel(logLevel)
		if err != nil {
			level = zerolog.DebugLevel // in case of error we default to debug
		}

		zerolog.SetGlobalLevel(level)
		output := os.Stdout
		writer := zerolog.New(output)

		if envLevel != "stg" && envLevel != "prd" && envLevel != "dev" {
			writer = zerolog.New(zerolog.ConsoleWriter{
				Out:        output,
				TimeFormat: time.RFC3339,
				FormatMessage: func(i any) string {
					return fmt.Sprintf("| %s |", i)
				},
			})
		}

		log.Logger = writer.With().
			Timestamp().
			Caller(). // Caller (file and line number) to every log event
			Str("env", envLevel).
			Str("service", config.Cfg.NAME).
			Int("pid", os.Getpid()).
			Str("lang", "Golang").
			Str("go_version", buildInfo.GoVersion).
			Logger()

		log.Error().Msg("🪵 Logger initialized!")
	})
}
