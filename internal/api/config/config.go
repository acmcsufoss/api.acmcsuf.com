package config

import "github.com/acmcsufoss/api.acmcsuf.com/utils"

type Config struct {
	Env            string
	Port           string
	DatabaseURL    string
	TrustedProxies []string
	AllowedOrigins []string
}

func Load() *Config {
	return &Config{
		Env:			utils.GetEnv("ENV", "development"),
		Port:           utils.GetEnv("PORT", "8080"),
		DatabaseURL:    utils.GetEnv("DATABASE_URL", "file:dev.db?cache=shared&mode=rwc"),
		TrustedProxies: utils.GetEnvAsSlice("TRUSTED_PROXIES", []string{"127.0.0.1/32"}),
		AllowedOrigins: utils.GetEnvAsSlice("ALLOWED_ORIGINS", []string{"*"}),
	}
}
