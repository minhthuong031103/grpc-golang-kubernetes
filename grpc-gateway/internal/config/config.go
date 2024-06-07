package config

// Configuration stores all the configuration data needed by the application
type Configuration struct {
	OrderServiceAddress string
	ServerAddress       string
}

// Load fetches configuration required for running the application
func Load() *Configuration {
	return &Configuration{
		OrderServiceAddress: "order-service:50051",
		ServerAddress:       "0.0.0.0:8080",
	}
}
