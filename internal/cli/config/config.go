package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	APIURL   string `json:"api_url"`
	LogLevel string `json:"log_level"`
}

var defaultConfig = Config{
	APIURL:   "http://localhost:8080",
	LogLevel: "info",
}

// Global config variable
var Cfg *Config

func init() {
	Cfg = &Config{}
}

// Loads config with three layers of precedence
// 1. Start with default config
// 2. Load values from config file if present
// 3. Apply the API URL override passed on the command line (if any)
func Load(apiURLOverride string) (*Config, error) {
	// Load default config
	cfg := &Config{
		APIURL:   defaultConfig.APIURL,
		LogLevel: defaultConfig.LogLevel,
	}

	// Override with stuff from config file (if present)
	path, err := getConfigPath()
	if err != nil {
		// TODO: Set to warning level when we have a better logger
		log.Printf("Warning: could not get config path. Reason: %v", err)
	} else { // Skips loading config from file if couldn't get path
		if data, err := os.ReadFile(path); err == nil {
			if err := json.Unmarshal(data, cfg); err != nil {
				log.Printf("Warning: failed to parse config file: %v", err)
			}
		}
	}

	if apiURLOverride != "" {
		cfg.APIURL = apiURLOverride
	}

	apiURL, err := NormalizeAPIURL(cfg.APIURL)
	if err != nil {
		return nil, err
	}
	cfg.APIURL = apiURL

	return cfg, nil
}

func getConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	// '~/.config/acmcsuf-cli/config.json' on Unix systems
	appDir := filepath.Join(configDir, "acmcsuf-cli")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(appDir, "config.json"), nil
}

func createDefaultConfigFile() error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	// Check if file exists and error out if it does (don't want to overwrite or truncate)
	if _, err := os.Stat(path); err == nil {
		return errors.New("config file already exists (refusing to overwrite)")
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")                           // pretty print
	if err := encoder.Encode(defaultConfig); err != nil { // writes to file here
		return err
	}
	return nil
}
