package ginmiddleware

import (
	"context"
	"net/http"
	"strings"

	customerpb "gateway/internal/generated/customer"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks the Authorization header for a valid token
func AuthMiddleware(authclient customerpb.CustomerServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Extract the token from the header
		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}

		// Validate the token (this is a placeholder, replace with your own validation logic)
		auth, err := authclient.Authorize(context.Background(), &customerpb.Token{
			Token: token,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("customer_id", auth.CustomerId)
		c.Set("email", auth.Email)
		c.Next()
	}
}

type NoAuthURL struct {
	Url    string
	Method string
}

// ConditionalAuthMiddleware applies AuthMiddleware conditionally
func ConditionalAuthMiddleware(noAuthURLs []NoAuthURL, authclient customerpb.CustomerServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		for _, u := range noAuthURLs {
			if c.Request.URL.Path == u.Url && c.Request.Method == u.Method {
				c.Next()
				return
			}
		}

		AuthMiddleware(authclient)(c)
	}
}
