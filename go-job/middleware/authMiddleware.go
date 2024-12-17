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

// func AuthMiddleware(c *gin.Context) {
// 	// Get the access token from cookies
// 	token, err := c.Cookie("access_token")
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not found"})
// 		c.Abort()
// 		return
// 	}

// 	// Validate the access token
// 	claims, err := utils.ValidateJWT(token, false)
// 	if err != nil {
// 		// If the access token is expired, try refreshing with the refresh token
// 		if err.Error() == "access token has expired" {
// 			// Get the refresh token from cookies
// 			refreshToken, err := c.Cookie("refresh_token")
// 			if err != nil {
// 				// If no refresh token is found, respond with unauthorized
// 				c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found"})
// 				c.Abort()
// 				return
// 			}

// 			// Validate the refresh token
// 			refreshClaims, err := utils.ValidateJWT(refreshToken, true)
// 			if err != nil {
// 				// If the refresh token is invalid, respond with unauthorized
// 				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
// 				c.Abort()
// 				return
// 			}

// 			// Generate new access token and refresh token
// 			accessToken, newRefreshToken, err := utils.GenerateJWT(refreshClaims.UserID, refreshClaims.Name, refreshClaims.Email)
// 			if err != nil {
// 				// If there's an error generating the tokens, respond with an internal server error
// 				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate new access token"})
// 				c.Abort()
// 				return
// 			}

// 			// Set the new tokens in cookies
// 			utils.SetCookies(c, accessToken, newRefreshToken)

// 			// Return the new tokens in the response
// 			c.JSON(http.StatusOK, gin.H{
// 				"access_token":  accessToken,
// 				"refresh_token": newRefreshToken,
// 			})
// 			// Stop further processing as we've already returned a response
// 			c.Abort()
// 			return
// 		}

// 		// If there's another error with the access token, respond with unauthorized
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid or expired token: %v", err)})
// 		c.Abort()
// 		return
// 	}

// 	// If token is valid, store user data in the context for downstream use
// 	c.Set("userID", claims.UserID)
// 	c.Set("name", claims.Name)
// 	c.Set("email", claims.Email)

// 	// Continue to the next handler
// 	c.Next()
// }
