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

	fileuploadclient := grpcclientconn.NewGRPCClient(cfg.FileUploadService.Address)
	if err := fileuploadclient.Connect(); err != nil {
		log.Fatalf("Failed to connect to file upload service: %v", err)
	}

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
	server := server.NewServer(fileuploadclient.GetConnection(), customerclient.GetConnection(), productclient.GetConnection(), orderclient.GetConnection())
	server.RegisterServiceHandler()
	server.SetRoute()
	server.Run(cfg.Server.Port)
}
