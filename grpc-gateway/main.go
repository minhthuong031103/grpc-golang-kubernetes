package main

import (
	"fmt"
	"gateway/config"
	grpcclientconn "gateway/internal/connection"
	"gateway/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	fmt.Printf("GATEWAY is trying serve on port: %d\n", cfg.Server.Port)

	orderclient := grpcclientconn.NewGRPCClient(cfg.OrderSvc.Address)
	if err := orderclient.Connect(); err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}

	defer orderclient.Disconnect()

	// Set up the HTTP server with integrated gRPC-Gateway and Gin router
	router := server.SetupRouter(orderclient.GetConnection())

	// Start the HTTP server and log any errors encountered
	log.Printf("API gateway server is running on %d", cfg.Server.Port)
	if err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatalf("gateway server closed abruptly: %v", err)
	}
}
