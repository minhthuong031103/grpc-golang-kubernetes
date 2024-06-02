// ./cmd/client/main.go

package main

import (
	"context"
	"log"
	"net/http"

	orderpb "grpc-gateway/protogen/golang/order"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Set up a connection to the order server.
	orderServiceAddr := "order-service:50051"
	conn, err := grpc.NewClient(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}
	defer conn.Close()

	// Create a gRPC-Gateway mux
	mux := runtime.NewServeMux()
	if err := orderpb.RegisterOrdersHandler(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the order server: %v", err)
	}

	// Create a Gin router
	r := gin.Default()

	// Integrate the gRPC-Gateway mux with Gin
	r.Any("/v0/*any", gin.WrapH(mux))

	// Add additional routes or middleware as needed
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Start the HTTP server
	addr := "0.0.0.0:8080"
	log.Printf("API gateway server is running on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}
}
