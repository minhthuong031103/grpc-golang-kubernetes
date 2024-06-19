package server

import (
	ginmiddleware "gateway/internal/middleware"
	"net/http"
	"strings"

	customerpb "gateway/internal/generated/customer"

	"github.com/gin-gonic/gin"
)

func (s *HTTPServer) SetRoute() {
	s.Router.Use(ginmiddleware.ConditionalAuthMiddleware([]ginmiddleware.NoAuthURL{
		{Url: "/api/login", Method: http.MethodPost},
		{Url: "/api/signup", Method: http.MethodPost},
		{Url: "/api/iam", Method: http.MethodPost},
	},
		customerpb.NewCustomerServiceClient(s.CustomerConn),
	))

	// File upload endpoint
	s.Router.POST("/upload", uploadFileHandler(s.FileUploadConn))

	// Handle all requests using gRPC-Gateway with a specific prefix, e.g., `/api`
	apiRoutes := s.Router.Group("/api")
	apiRoutes.Any("/*any", func(ctx *gin.Context) {
		ctx.Request.URL.Path = strings.TrimPrefix(ctx.Request.URL.Path, "/api")
		ctx.Next()
	}, gin.WrapH(s.Gwmux))
}
