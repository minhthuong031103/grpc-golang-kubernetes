package server

import (
	"fmt"
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
	}, func(c *gin.Context) {
		// Check permission for specific endpoints

		role, existRole := c.Get("role")
		if !existRole {
			fmt.Println("Role not found")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		if role == "customer" {
			type AllowEndpoint struct {
				Method string
				Path   string
			}

			allowEndpoints := []AllowEndpoint{
				{Method: http.MethodGet, Path: "/products"},
				{Method: http.MethodGet, Path: "/product/*"},
				{Method: http.MethodPost, Path: "/orders"},
				{Method: http.MethodGet, Path: "/orders"},
				{Method: http.MethodGet, Path: "/orders"},
			}

			isAllowed := false
			for _, endpoint := range allowEndpoints {
				fmt.Println(c.Request.Method, endpoint.Method, c.Request.URL.Path, endpoint.Path)
				fmt.Println(endpoint.Path[:len(endpoint.Path)-2])
				if c.Request.Method == endpoint.Method && (c.Request.URL.Path == endpoint.Path || (strings.HasSuffix(endpoint.Path, "/*") && strings.HasPrefix(c.Request.URL.Path, endpoint.Path[:len(endpoint.Path)-2]))) {
					isAllowed = true
					break
				}
			}

			if !isAllowed {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
				return
			}
		}

		// Forward the request to the gRPC-Gateway
		c.Next()
	}, gin.WrapH(s.Gwmux))
}
