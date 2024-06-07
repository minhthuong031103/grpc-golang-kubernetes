package server

import (
	"context"
	"log"
	"net/http"

	orderpb "grpc-gateway/protogen/golang/order"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// SetupRouter initializes and returns a new Gin router
func SetupRouter(grpcConn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	// Set up the gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	err := orderpb.RegisterOrdersHandler(context.Background(), gwmux, grpcConn)
	if err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}

	// Handle all requests using gRPC-Gateway with a specific prefix, e.g., `/api`
	apiRoutes := router.Group("/api")
	apiRoutes.Any("/*any", gin.WrapH(gwmux))

	// Standard HTTP routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Additional routes can be added here
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return router
}
