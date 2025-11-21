package config

import "github.com/acmcsufoss/api.acmcsuf.com/utils"

type Config struct {
	Port           string
	DatabaseURL    string
	AllowedOrigins []string
}

func Load() *Config {
	return &Config{
		Port:           utils.GetEnv("PORT", "8080"),
		AllowedOrigins: utils.GetEnvAsSlice("ALLOWED_ORIGINS", []string{"*"}),
		DatabaseURL:    utils.GetEnv("DATABASE_URL", "file:dev.db?cache=shared&mode=rwc"),
	}
}
