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

	customerclient := grpcclientconn.NewGRPCClient(cfg.CustomerSvc.Address)
	if err := customerclient.Connect(); err != nil {
		log.Fatalf("Failed to connect to customer service: %v", err)
	}
	log.Println("Connected to CUSTOMER service")
	defer customerclient.Disconnect()

	productclient := grpcclientconn.NewGRPCClient(cfg.ProductSvc.Address)
	if err := productclient.Connect(); err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	log.Println("Connected to PRODUCT service")
	defer productclient.Disconnect()

	orderclient := grpcclientconn.NewGRPCClient(cfg.OrderSvc.Address)
	if err := orderclient.Connect(); err != nil {
		log.Fatalf("Failed to connect to order service: %v", err)
	}
	log.Println("Connected to ORDER service")
	defer orderclient.Disconnect()

	// Set up the HTTP server with integrated gRPC-Gateway and Gin router
	router := server.SetupRouter(customerclient.GetConnection(), productclient.GetConnection(), orderclient.GetConnection())

	fmt.Println("CUSTOMER SERVICE ADDRESS: ", cfg.CustomerSvc.Address)

	// Start the HTTP server and log any errors encountered
	log.Printf("API gateway server is running on %d", cfg.Server.Port)
	if err := router.Run(fmt.Sprintf(":%d", cfg.Server.Port)); err != nil {
		log.Fatalf("gateway server closed abruptly: %v", err)
	}
}
