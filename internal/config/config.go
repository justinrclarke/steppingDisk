package config

import (
	"encoding/json"
	"os"
)

type DatabaseConfig struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	ListenAddr    string           `json:"listen_addr"`
	Databases     []DatabaseConfig `json:"databases"`
	EncryptionKey string           `json:"encryption_key"`
}

func Load() (*Config, error) {
	f, err := os.Open("config/config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
