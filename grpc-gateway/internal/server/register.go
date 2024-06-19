package server

import (
	"context"
	customerpb "gateway/internal/generated/customer"
	fileuploadpb "gateway/internal/generated/fileupload"
	orderpb "gateway/internal/generated/order"
	productpb "gateway/internal/generated/product"
	"log"
)

func (s *HTTPServer) RegisterServiceHandler() {
	err := fileuploadpb.RegisterFileUploadServiceHandler(context.Background(), s.Gwmux, s.FileUploadConn)
	if err != nil {
		log.Fatalf("Failed to register FILEUPLOAD: %v", err)
	}

	err = customerpb.RegisterCustomerServiceHandler(context.Background(), s.Gwmux, s.CustomerConn)
	if err != nil {
		log.Fatalf("Failed to register CUSTOMER: %v", err)
	}

	err = productpb.RegisterProductServiceHandler(context.Background(), s.Gwmux, s.ProductConn)
	if err != nil {
		log.Fatalf("Failed to register PRODUCT: %v", err)
	}

	err = orderpb.RegisterOrderServiceHandler(context.Background(), s.Gwmux, s.OrderConn)
	if err != nil {
		log.Fatalf("Failed to register ORDER: %v", err)
	}

}
