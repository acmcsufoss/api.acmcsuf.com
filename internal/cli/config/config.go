package main

// package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
func Load() (*Config, error) {
	cfg := &defaultConfig
	var err error
	path, err := getConfigPath()
	if err != nil {
		return nil, err
	}

	var file *os.File
	file, err = os.Open(path)
	if err != nil {
		// If file doesn't exist, create one with the default config and open it
		if errors.Is(err, os.ErrNotExist) {
			err = createDefaultConfigFile()
			if err != nil {
				panic(err)
			}
			// Try to open file again (shuold be created now)
			file, err = os.Open(path)
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
	defer file.Close()

	var config *Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
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
	path, _ := getConfigPath()
	fmt.Printf("path: '%s'\n", path)
	_ = createDefaultConfigFile()
}
