package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve session ID from request header
		sessionID := c.GetHeader("Authorization")

		// Check if session ID is empty
		if sessionID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Perform session validation (dummy implementation)
		if !isValidSession(sessionID) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		// Continue with next middleware/handler if session is valid
		c.Next()
	}
}

// Dummy session validation function (replace with actual implementation)
func isValidSession(sessionID string) bool {
	// In a real-world scenario, this function should check if the session is valid
	// For example, you could check if the session exists in a session store
	// or validate the session using a JWT library
	// For demonstration purposes, we'll just check if the sessionID is non-empty
	return sessionID != ""
}
