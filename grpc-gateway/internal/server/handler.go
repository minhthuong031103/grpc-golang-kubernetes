package server

import (
	"context"
	customerpb "gateway/internal/generated/customer"
	fileuploadpb "gateway/internal/generated/fileupload"
	orderpb "gateway/internal/generated/order"
	productpb "gateway/internal/generated/product"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func (s *HTTPServer) SetGmux() {

	gwmux := runtime.NewServeMux()

	err := fileuploadpb.RegisterFileUploadServiceHandler(context.Background(), gwmux, s.FileUploadConn)
	if err != nil {
		log.Fatalf("Failed to register FILEUPLOAD: %v", err)
	}

	err = customerpb.RegisterCustomerServiceHandler(context.Background(), gwmux, s.CustomerConn)
	if err != nil {
		log.Fatalf("Failed to register CUSTOMER: %v", err)
	}

	err = productpb.RegisterProductServiceHandler(context.Background(), gwmux, s.ProductConn)
	if err != nil {
		log.Fatalf("Failed to register PRODUCT: %v", err)
	}

	err = orderpb.RegisterOrderServiceHandler(context.Background(), gwmux, s.OrderConn)
	if err != nil {
		log.Fatalf("Failed to register ORDER: %v", err)
	}

	// File upload endpoint
	s.Router.POST("/upload", uploadFileHandler(s.FileUploadConn))

	// Handle all requests using gRPC-Gateway with a specific prefix, e.g., `/api`
	apiRoutes := s.Router.Group("/api")
	apiRoutes.Any("/*any", func(ctx *gin.Context) {
		ctx.Request.URL.Path = strings.TrimPrefix(ctx.Request.URL.Path, "/api")
		ctx.Next()
	}, gin.WrapH(gwmux))
}
