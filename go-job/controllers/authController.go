package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/mysterybee07/go-react-job/database"
	"github.com/mysterybee07/go-react-job/models"
	"github.com/mysterybee07/go-react-job/payloads"
	"github.com/mysterybee07/go-react-job/utils"
)

func RegisterUser(c *gin.Context) {
	var input payloads.RegisterUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(
			"Failed to parse JSON body. Details: %s", err.Error(),
		)})
		return
	}
	// fmt.Println(len(input.Password))

	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", input.ContactEmail, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to process password",
		})
		return
	}

	input.Password = hashPassword

	user := models.User{
		ContactInfo: models.ContactInfo{
			Name:         input.Name,
			ContactEmail: input.ContactEmail,
			ContactPhone: input.ContactPhone,
			ImageUrl:     input.ImageUrl,
			Address:      input.Address,
		},
		Resume:   input.Resume,
		Password: input.Password,
	}

	var existingUser models.User
	if err := database.DB.Where("contact_email = ? OR contact_phone = ?", user.ContactEmail, user.ContactPhone).First(&existingUser).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or phone number already exists.",
		})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create user. Please try again.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "User created successfully",
	})
}

func RegisterCompany(c *gin.Context) {
	var input payloads.RegisterCompany

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error parsing data. Details: %v", err.Error()),
		})
		return
	}

	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", input.ContactEmail, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to process password",
		})
		return
	}
	input.Password = hashPassword
	company := models.Company{
		ContactInfo: models.ContactInfo{
			Name:         input.Name,
			ContactEmail: input.ContactEmail,
			ContactPhone: input.ContactPhone,
			ImageUrl:     input.ImageUrl,
			Address:      input.Address,
		},
		Description: input.Description,
		Password:    input.Password,
	}
	if err := database.DB.Create(&company).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to create company. Details: %v", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "company created successfully",
		"company": company,
	})
}

func Login(c *gin.Context) {
	type LoginData struct {
		ContactEmail string `json:"contact_email"`
		Password     string `json:"password"`
	}
	var loginData LoginData

	// Bind JSON to struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error parsing data. Details: %v", err.Error()),
		})
		return
	}

	// Check for the email in User table
	var user models.User
	err := database.DB.Where("contact_email = ?", loginData.ContactEmail).First(&user).Error

	if err != nil {
		// If not found in User, attempt Company login
		token, refreshToken, err := utils.CompanyLogin(loginData.ContactEmail, loginData.Password)
		if err != nil {
			if err.Error() == "invalid password" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"error": "Incorrect password.",
				})
			} else {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Email not found.",
				})
			}
			return
		}

		// Set cookies for company
		utils.SetCookies(c, token, refreshToken)

		// Send response for company
		c.JSON(http.StatusOK, gin.H{
			// "access_token":  token,
			// "refresh_token": refreshToken,
			"user_type": "company",
			"message":   "Login successful",
		})
		return
	}

	// Validate password for user
	match := utils.CheckPasswordHash(loginData.Password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Incorrect password.",
		})
		return
	}

	// Generate JWT tokens for user
	userID := strconv.Itoa(int(user.ID))
	token, refreshToken, err := utils.GenerateJWT(userID, user.Name, user.ContactEmail)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error generating tokens: %v", err.Error()),
		})
		return
	}

	// Set cookies for user
	utils.SetCookies(c, token, refreshToken)

	// Send response for user
	c.JSON(http.StatusOK, gin.H{
		// "access_token":  token,
		// "refresh_token": refreshToken,
		"user_type": "user",
		"message":   "Login successful",
	})
}

func AuthorizedUser(c *gin.Context) {
	// Retrieve user ID from the context
	userID, exists := c.Get("userID")
	fmt.Println(userID)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "User ID not found in context",
		})
		return
	}

	// Retrieve user information from database
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		log.Println("User not found in database")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	// Return logged-in user's details
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"ID":    user.ID,
			"email": user.ContactEmail,
		},
		"message": "User data retrieved successfully",
	})
}

// Logout clears the access and refresh token cookies
func Logout(c *gin.Context) {

	c.SetCookie(
		"access_token",
		"", // Set empty value
		-1, // Expiry time (negative means immediate removal)
		"/",
		"",
		false,
		true,
	)

	c.SetCookie(
		"refresh_token",
		"",
		-1,
		"/",
		"",
		false,
		true,
	)

	// Send a successful response
	c.JSON(http.StatusOK, gin.H{
		"message":   "Logout successful",
		"timestamp": time.Now(),
	})
}

func RefreshToken(c *gin.Context) {
	// Retrieve the refresh token from cookies
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Refresh token missing or invalid",
		})
		return
	}

	// Validate the refresh token
	claims, err := utils.ValidateJWT(refreshToken, true)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid refresh token: " + err.Error(),
		})
		return
	}

	// Generate a new access token (reusing user details from the refresh token)
	newAccessToken, _, err := utils.GenerateJWT(claims.UserID, claims.Name, claims.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate new access token",
		})
		return
	}

	// Update the access token in cookies
	utils.SetCookies(c, newAccessToken, refreshToken)

	// Send the new access token as a response
	c.JSON(http.StatusOK, gin.H{
		"access_token": newAccessToken,
		"message":      "New access token generated successfully",
	})
}
