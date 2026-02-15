package config

import (
	"fmt"
	"net/url"
)

// Intended to be used with .JoinPath() to construct URLs with configured host:port
func GetBaseURL(cfg *Config) *url.URL {
	return &url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	}
}
