package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

//go:embed config.yaml
var configFile []byte

type Config struct {
	Server            ServerConfig  `yaml:"server"`
	FileUploadService ServiceConfig `yaml:"fileuploadsvc"`
	CustomerSvc       ServiceConfig `yaml:"customersvc"`
	OrderSvc          ServiceConfig `yaml:"ordersvc"`
	ProductSvc        ServiceConfig `yaml:"productsvc"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type ServiceConfig struct {
	Address string `yaml:"address"`
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
}

func LoadConfig() (*Config, error) {
	var config Config
	err := yaml.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config data: %v", err)
	}

	config.FileUploadService.Address = fmt.Sprintf("%s:%d", config.FileUploadService.Host, config.FileUploadService.Port)
	config.CustomerSvc.Address = fmt.Sprintf("%s:%d", config.CustomerSvc.Host, config.CustomerSvc.Port)
	config.OrderSvc.Address = fmt.Sprintf("%s:%d", config.OrderSvc.Host, config.OrderSvc.Port)
	config.ProductSvc.Address = fmt.Sprintf("%s:%d", config.ProductSvc.Host, config.ProductSvc.Port)
	return &config, nil
}
