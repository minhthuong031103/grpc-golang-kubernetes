package server

import (
	"context"
	"log"
	"net/http"
	"strings"

	orderpb "gateway/internal/generated/order"
	productpb "gateway/internal/generated/product"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// SetupRouter initializes and returns a new Gin router
func SetupRouter(productConn *grpc.ClientConn, orderConn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()

	// Set up the gRPC-Gateway mux
	gwmux := runtime.NewServeMux()
	err := productpb.RegisterProductServiceHandler(context.Background(), gwmux, productConn)
	if err != nil {
		log.Fatalf("Failed to register PRODUCT: %v", err)
	}
	err = orderpb.RegisterOrdersHandler(context.Background(), gwmux, orderConn)
	if err != nil {
		log.Fatalf("Failed to register ORDER: %v", err)
	}

	// Standard HTTP routes
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Additional routes can be added here
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// Handle all requests using gRPC-Gateway with a specific prefix, e.g., `/api`
	apiRoutes := router.Group("/api")
	apiRoutes.Any("/*any", func(ctx *gin.Context) {
		ctx.Request.URL.Path = strings.TrimPrefix(ctx.Request.URL.Path, "/api")
		ctx.Next()
	}, gin.WrapH(gwmux))

	return router
}
