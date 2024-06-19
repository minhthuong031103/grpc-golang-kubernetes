package server

import (
	"fmt"
	"log"
	"net/http"

	ginmiddleware "gateway/middleware"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type HTTPServer struct {
	Router         *gin.Engine
	FileUploadConn *grpc.ClientConn
	CustomerConn   *grpc.ClientConn
	ProductConn    *grpc.ClientConn
	OrderConn      *grpc.ClientConn
}

func (s *HTTPServer) Run(port int) {
	log.Printf("API gateway server is running on %d", port)
	if err := s.Router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalf("gateway server closed abruptly: %v", err)
	}
}

func NewServer(fileuploadConn *grpc.ClientConn, customerConn *grpc.ClientConn, productConn *grpc.ClientConn, orderConn *grpc.ClientConn) *HTTPServer {
	router := gin.Default()
	router.Use(ginmiddleware.GetCORSMiddleware())
	router.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return &HTTPServer{
		Router:         router,
		FileUploadConn: fileuploadConn,
		CustomerConn:   customerConn,
		ProductConn:    productConn,
		OrderConn:      orderConn,
	}
}
