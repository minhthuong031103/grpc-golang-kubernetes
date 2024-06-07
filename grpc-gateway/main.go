package main

import (
	"log"

	"grpc-gateway/internal/config"
	"grpc-gateway/internal/orderclient"
	"grpc-gateway/internal/server"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Establish a connection to the order service
	orderclient := orderclient.NewOrderClient(cfg.OrderServiceAddress)
	if err := orderclient.Connect(); err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	defer orderclient.Disconnect()

	// Set up the HTTP server with integrated gRPC-Gateway and Gin router
	router := server.SetupRouter(orderclient.GetConnection())

	// Start the HTTP server and log any errors encountered
	log.Printf("API gateway server is running on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("gateway server closed abruptly: %v", err)
	}
}
