package main

import (
	"fmt"

	"github.com/gsn_budget_service/internal/config"
	"github.com/gsn_budget_service/pkg/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	config.LoadConfig()
	logger.InitLogger()

	log.Info().Msg("Dude")
	fmt.Println("Hello World, Gustavo")
}
