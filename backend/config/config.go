package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

type Config struct {
	Minio    `json:"minio"`
	Database `json:"database"`
}

// Database struct
type Database struct {
	DBHost   string `json:"db_host"`
	DBPort   string `json:"db_port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

// Minio struct
type Minio struct {
	Host      string `json:"host"`
	AccessKey string `json:"access_key"`
	SecretKey string `json:"secret_key"`
}

// Return New Config
func NewConfig() *Config {
	return &Config{}
}

// LoadConfig method
// read file and parse into config file
func (c *Config) LoadConfig(path string) (*Config, error) {
	config := new(Config)
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
	fmt.Println(config)
	return config, nil
}
