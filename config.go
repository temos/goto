package main

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Directories []ConfigDirectory `yaml:"directories"`
	URLs        []ConfigURL       `yaml:"urls"`
	ActiveColor string            `yaml:"activeColor"`
}

type ConfigDirectory struct {
	Prefix string `yaml:"prefix"`
	Path   string `yaml:"path"`

	ShowHidden bool `yaml:"showHidden"`
}

type ConfigURL struct {
	Prefix string `yaml:"prefix"`
	Name   string `yaml:"name"`
	URL    string `yaml:"url"`
}

// LoadConfig loads the configuration from a yaml file at the provided path
func LoadConfig(path string) (*Config, error) {
	configBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read config: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	if config.ActiveColor == "" {
		config.ActiveColor = "#8C18E2"
	}

	return &config, nil
}
