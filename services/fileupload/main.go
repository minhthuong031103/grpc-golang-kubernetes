package main

import (
	"log"

	"fileuploadsvc/config"
	"fileuploadsvc/internal/server"
)

func main() {
	log.Println("Starting FILE-UPLOADER service...")
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server.StartGRPCServer(config.Server.Port, config.S3Storage)
}
