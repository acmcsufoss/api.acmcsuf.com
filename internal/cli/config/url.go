package config

import (
	"fmt"
	"net/url"
	"strings"
)

// Intended to be used with .JoinPath() to construct URLs
func GetBaseURL(cfg *Config) *url.URL {
	apiURL, _ := url.Parse(cfg.APIURL)
	return apiURL
}

func NormalizeAPIURL(rawURL string) (string, error) {
	apiURL, err := url.Parse(strings.TrimSpace(rawURL))
	if err != nil {
		return "", fmt.Errorf("invalid API URL: %w", err)
	}
	if apiURL.Scheme != "http" && apiURL.Scheme != "https" {
		return "", fmt.Errorf("API URL must use http or https")
	}
	if apiURL.Host == "" {
		return "", fmt.Errorf("API URL must include a host")
	}
	if apiURL.User != nil || apiURL.RawQuery != "" || apiURL.Fragment != "" {
		return "", fmt.Errorf("API URL must not include credentials, a query, or a fragment")
	}
	if apiURL.Path != "" && apiURL.Path != "/" {
		return "", fmt.Errorf("API URL must not include a path")
	}

	apiURL.Path = ""
	return apiURL.String(), nil
}
