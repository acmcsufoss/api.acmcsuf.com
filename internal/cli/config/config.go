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
	ServerURI string `json:"server_uri"`
}

func Load() (*Config, error) {
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
			err = createDefaultConfig()
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

func createDefaultConfig() error {
	path, err := getConfigPath()
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	// Default settings ==========================================
	defaultConfig := &Config{
		ServerURI: "http://localhost:8080",
	}
	// ===========================================================
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
	_ = createDefaultConfig()
}
