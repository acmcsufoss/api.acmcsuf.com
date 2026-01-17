package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
}

// Subset of Config struct that can be overridden with command line flags
type ConfigOverrides struct {
	Host string
	Port string
}

var defaultConfig = Config{
	Host:     "localhost",
	Port:     "8080",
	LogLevel: "info",
}

// Global config variable
var Cfg *Config

// Loads config with three layers of precedence
// 1. Start with default config
// 2. Load values from config file if present
// 3. Provide any overrides passed in through command line flags (if any)
func Load(overrides *ConfigOverrides) (*Config, error) {
	// Load default config
	cfg := &Config{
		Host:     defaultConfig.Host,
		Port:     defaultConfig.Port,
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
