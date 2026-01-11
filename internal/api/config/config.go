package config

import (
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

type Config struct {
	Env             string
	Port            string
	DatabaseURL     string
	TrustedProxies  []string
	AllowedOrigins  []string
	GuildID         string
	DiscordBotToken string
}

func Load() *Config {
	return &Config{
		// Non-sensistive
		Env:            utils.GetEnv("ENV", "development"),
		Port:           utils.GetEnv("PORT", "8080"),
		DatabaseURL:    utils.GetEnv("DATABASE_URL", "file:dev.db?cache=shared&mode=rwc"),
		TrustedProxies: utils.GetEnvAsSlice("TRUSTED_PROXIES", []string{"127.0.0.1/32"}),
		AllowedOrigins: utils.GetEnvAsSlice("ALLOWED_ORIGINS", []string{"http://acmcsuf-api-dev"}),
		GuildID:        utils.GetEnv("GUILD_ID", "710225099923521558"), // acmcsuf's GuildID

		// Sensitive (no default)
		DiscordBotToken: os.Getenv("DISCORD_BOT_TOKEN"),
	}
}
