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
	DB_URL      string
	DB_HOST     string
	DB_PORT     uint32
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_SSL_MODE string
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

		// APP Settings
		viper.SetDefault("NAME", "gsn_budget_service")
		viper.SetDefault("ENVIRONMENT", env)
		viper.SetDefault("LOG_LEVEL", "DEBUG")
		viper.SetDefault("PORT", 8080)
		viper.SetDefault("URL_PREFIX", "budget_api")

		// ENTRYPOINTS
		viper.SetDefault("API_URL", "http://localhost:8080")

		// DB
		viper.SetDefault("DB_HOST", "localhost")
		viper.SetDefault("DB_PORT", 5432)
		viper.SetDefault("DB_USER", "postgres")
		viper.SetDefault("DB_PASSWORD", "password")
		viper.SetDefault("DB_NAME", "")
		viper.SetDefault("DB_SSL_MODE", "disable")

		name := viper.GetString("NAME")
		environment := viper.GetString("ENVIRONMENT")
		logLevel := viper.GetString("LOG_LEVEL")
		port := viper.GetInt32("PORT")
		urlPrefix := viper.GetString("URL_PREFIX")
		apiUrl := viper.GetString("API_URL")

		dbHost := viper.GetString("DB_HOST")
		dbPort := viper.GetUint32("DB_PORT")
		dbUser := viper.GetString("DB_USER")
		dbPassword := viper.GetString("DB_PASSWORD")
		dbName := viper.GetString("DB_NAME")
		dbSslMode := viper.GetString("DB_SSL_MODE")

		dbUrl := fmt.Sprintf(
			"postgres://%s:%s@%s:%d/%s?sslmode=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSslMode,
		)

		viper.SetDefault("DB_URL", dbUrl)

		Cfg = &Config{
			NAME:        name,
			ENVIRONMENT: environment,
			LOG_LEVEL:   logLevel,
			PORT:        port,
			URL_PREFIX:  urlPrefix,
			API_URL:     apiUrl,
			DB_URL:      dbUrl,
			DB_HOST:     dbHost,
			DB_PORT:     dbPort,
			DB_USER:     dbUser,
			DB_PASSWORD: dbPassword,
			DB_NAME:     dbName,
			DB_SSL_MODE: dbSslMode,
		}
	})

	return Cfg
}
