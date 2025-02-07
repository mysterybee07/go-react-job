package utils

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	Email  string
	Name   string
	UserID string
	jwt.StandardClaims
}

func GenerateJWT(userID string, name, email string) (string, string, error) {
	expirationTime := time.Now().UTC().Add(5 * time.Minute)
	claims := &SignedDetails{
		UserID: userID,
		Name:   name,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(7 * 24 * time.Hour).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtSecret)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign token: %w", err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(jwtSecret)
	if err != nil {
		return "", "", fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return token, refreshToken, nil
}

// SetCookies sets the access and refresh tokens as HTTP-only cookies
func SetCookies(c *gin.Context, accessToken string, refreshToken string) {
	// Access token cookie
	c.SetCookie(
		"access_token", // Cookie name
		accessToken,    // Cookie value
		900,            // MaxAge in seconds (15 minutes)
		"/",            // Path
		"",             // Domain (optional, leave blank for localhost)
		false,          // Secure: false for local development
		true,           // HttpOnly: true for security
	)

	// Refresh token cookie
	c.SetCookie(
		"refresh_token",
		refreshToken,
		604800,
		"/",
		"",
		false,
		true,
	)
}

func ValidateJWT(signedToken string, isRefreshToken bool) (*SignedDetails, error) {
	// Parse the signed token with claims
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("unable to parse token: %w", err)
	}

	// Extract the claims from the token
	claims, ok := token.Claims.(*SignedDetails)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// If this is a refresh token, ensure that it is valid
	if isRefreshToken {
		// You could also check for refresh token expiration
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("refresh token has expired")
		}
	} else {
		// Check if the access token has expired
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, fmt.Errorf("access token has expired")
		}
	}

	return claims, nil
}
