package ginmiddleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks the Authorization header for a valid token
func AuthMiddleware() gin.HandlerFunc {
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
		if !validateToken(token) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Token is valid, proceed to the next handler
		c.Next()
	}
}

// Placeholder function to validate the token
func validateToken(token string) bool {
	// Implement your token validation logic here
	// For example, you could decode a JWT and check its claims, expiration, etc.
	// This function should return true if the token is valid, false otherwise.
	return token == "your-valid-token" // Replace with real validation
}

type NoAuthURL struct {
	Url    string
	Method string
}

// ConditionalAuthMiddleware applies AuthMiddleware conditionally
func ConditionalAuthMiddleware(noAuthURLs []NoAuthURL) gin.HandlerFunc {
	return func(c *gin.Context) {

		for _, u := range noAuthURLs {
			if c.Request.URL.Path == u.Url && c.Request.Method == u.Method {
				// Skip auth middleware for these paths
				c.Next()
				return
			}
		}

		AuthMiddleware()(c)
	}
}
