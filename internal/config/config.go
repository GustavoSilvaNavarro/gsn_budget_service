package config

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type Config struct {
	// App details
	NAME        string
	ENVIRONMENT string
	LOG_LEVEL   string
	PORT        int32
	URL_PREFIX  string
	API_URL     string
}

var (
	Cfg  *Config
	once sync.Once
)

func LoadConfig() *Config {
	once.Do(func() {
		viper.AutomaticEnv()

		env := strings.ToLower(viper.GetString("ENVIRONMENT"))
		if env == "" {
			env = "local"
		}

		if env != "dev" && env != "stg" && env != "prd" {
			viper.SetConfigFile(".env")
			viper.SetConfigType("env")

			if err := viper.ReadInConfig(); err != nil {
				fmt.Println("⚠️ No secrets or env file has been found. Using default envs.")
			}
		}

		viper.SetDefault("NAME", "gsn_budget_service")
		viper.SetDefault("ENVIRONMENT", env)
		viper.SetDefault("LOG_LEVEL", "DEBUG")
		viper.SetDefault("PORT", 8080)
		viper.SetDefault("URL_PREFIX", "budget_api")
		viper.SetDefault("API_URL", "http://localhost:8080")

		Cfg = &Config{
			NAME:        viper.GetString("NAME"),
			ENVIRONMENT: viper.GetString("ENVIRONMENT"),
			LOG_LEVEL:   viper.GetString("LOG_LEVEL"),
			PORT:        viper.GetInt32("PORT"),
			URL_PREFIX:  viper.GetString("URL_PREFIX"),
			API_URL:     viper.GetString("API_URL"),
		}
	})

	return Cfg
}
