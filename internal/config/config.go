package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Hostname string `json:"hostname"`
	Port     uint16 `json:"port"`
	InMemory bool   `json:"inMemory"`
}

func Default() *Config {
	return &Config{
		Hostname: "localhost",
		Port:     48763,
		InMemory: true,
	}
}

func From(p string) (*Config, error) {
	config := &Config{}
	buf, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(buf, config); err != nil {
		return nil, err
	}

	return config, nil
}
