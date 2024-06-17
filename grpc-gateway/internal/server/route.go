package server

import (
	"context"
	"log"
	"net/http"
	"strings"

	customerpb "gateway/internal/generated/customer"
	fileuploadpb "gateway/internal/generated/fileupload"
	orderpb "gateway/internal/generated/order"
	productpb "gateway/internal/generated/product"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// SetupRouter initializes and returns a new Gin router
func SetupRouter(fileuploadConn *grpc.ClientConn, customerConn *grpc.ClientConn, productConn *grpc.ClientConn, orderConn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	// Set up CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))
	// Set up the gRPC-Gateway mux
	gwmux := runtime.NewServeMux()

	err := fileuploadpb.RegisterFileUploadServiceHandler(context.Background(), gwmux, fileuploadConn)
	if err != nil {
		log.Fatalf("Failed to register FILEUPLOAD: %v", err)
	}

	err = customerpb.RegisterCustomerServiceHandler(context.Background(), gwmux, customerConn)
	if err != nil {
		log.Fatalf("Failed to register CUSTOMER: %v", err)
	}

	err = productpb.RegisterProductServiceHandler(context.Background(), gwmux, productConn)
	if err != nil {
		log.Fatalf("Failed to register PRODUCT: %v", err)
	}

	err = orderpb.RegisterOrdersHandler(context.Background(), gwmux, orderConn)
	if err != nil {
		log.Fatalf("Failed to register ORDER: %v", err)
	}

	// Additional routes can be added here
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// File upload endpoint
	router.POST("/upload", uploadFileHandler(fileuploadConn))

	// Handle all requests using gRPC-Gateway with a specific prefix, e.g., `/api`
	apiRoutes := router.Group("/api")
	apiRoutes.Any("/*any", func(ctx *gin.Context) {
		ctx.Request.URL.Path = strings.TrimPrefix(ctx.Request.URL.Path, "/api")
		ctx.Next()
	}, gin.WrapH(gwmux))
	return router
}
