package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	"github.com/acmcsufoss/api.acmcsuf.com/utils"
)

type ConfigSource string

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
}

var defaultConfig = Config{
	Host:     "localhost",
	Port:     "8080",
	LogLevel: "info",
}

// Loads config with three layers of precedence
// 1. Start with default config
// 2. Load values from config file if present
// 3. Provide any overrides passed in thruogh command line flags (if any)
func Load(overrides *Config) (*Config, error) {
	// Load default config
	cfg := &defaultConfig

	// Override with stuff from config file (if present)
	path, err := getConfigPath()
	if err != nil {
		// TODO: Set to warning level when we have a better logger
		log.Printf("Warning: could not get config path. Reason: %v", err)
	} else { // Skips loading config from file if couldn't get path
		if data, err := os.ReadFile(path); err == nil {
			json.Unmarshal(data, cfg)
		}
	}

	// Override with args passed with command line args
	if overrides.Host != "" {
		cfg.Host = overrides.Host
	}
	if overrides.Port != "" {
		cfg.Port = overrides.Port
	}

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
	// TODO: This will go being a `acmcsuf-cli init` subcommand or something similar
	path, err := getConfigPath()
	if err != nil {
		return err
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

func main() {
	// createDefaultConfigFile()
	cfg, _ := Load()
	body, _ := json.Marshal(cfg)
	utils.PrettyPrintJSON(body)
}
