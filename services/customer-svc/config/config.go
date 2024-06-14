package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configFile []byte

type Config struct {
	Server    ServerConfig    `yaml:"server"`
	Cassandra CassandraConfig `yaml:"cassandra"`
	SecretKey string          `yaml:"secretkey"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type CassandraConfig struct {
	Hosts       []string `yaml:"hosts"`
	Keyspace    string   `yaml:"keyspace"`
	Consistency string   `yaml:"consistency"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %v", err)
	}
	return &config, nil
}
