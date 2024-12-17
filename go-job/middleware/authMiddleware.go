package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/utils"
)

// AuthMiddleware validates the access token and ensures the user is authorized
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the access token from cookies
		accessToken, err := c.Cookie("access_token")
		if err != nil || accessToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access token missing or invalid"})
			c.Abort()
			return
		}

		// Validate the access token
		claims, err := utils.ValidateJWT(accessToken, false)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: " + err.Error()})
			c.Abort()
			return
		}

		// Store claims in context for use in downstream handlers
		c.Set("userID", claims.UserID)
		c.Set("name", claims.Name)
		c.Set("email", claims.Email)

		// Pass control to the next handler
		c.Next()
	}
}
