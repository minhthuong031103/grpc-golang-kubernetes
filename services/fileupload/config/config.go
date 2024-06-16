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
	S3Storage S3StorageConfig `yaml:"s3storage"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}
type S3StorageConfig struct {
	Bucket    string `yaml:"bucket"`
	Region    string `yaml:"region"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %v", err)
	}
	return &config, nil
}
