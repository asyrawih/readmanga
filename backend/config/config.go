package config

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Minio
	Database
}

// Database struct
type Database struct {
	DBHost  string `json:"db_host,omitempty"`
	DBPort  string `json:"db_port,omitempty"`
	User    string `json:"user,omitempty"`
	Passwod string `json:"passwod,omitempty"`
	DBName  string `json:"db_name,omitempty"`
}

// Minio struct
type Minio struct {
	Host      string `json:"host,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key,omitempty"`
}

// Return New Config
func NewConfig() *Config {
	return &Config{}
}

// LoadConfig method
// read file and parse into config file
func (c *Config) LoadConfig(path string) (*Config, error) {
	var config *Config
	configByte, err := os.ReadFile(path)
	if err != nil {
		log.Err(err).Msg("")
		return nil, err
	}

	r := strings.NewReader(string(configByte))
	decoder := json.NewDecoder(r)
	// Decode into config
	if err := decoder.Decode(&config); err != nil {
		log.Err(err).Msg("")
		return nil, err
	}
	return config, nil
}
