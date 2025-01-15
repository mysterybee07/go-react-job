package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	// Bind and validate input
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to parse JSON body: %s", err.Error())})
		return
	}

	// Hash password
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", input.ContactEmail, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to process password"})
		return
	}
	input.Password = hashPassword

	// Create user object
	user := models.User{
		ContactInfo: models.ContactInfo{
			Name:         input.Name,
			ContactEmail: input.ContactEmail,
			ContactPhone: input.ContactPhone,
			Address:      input.Address,
		},
		Password: input.Password,
	}

	// Check for existing user
	var existingUser models.User
	if err := database.DB.Where("contact_email = ? OR contact_phone = ?", user.ContactEmail, user.ContactPhone).First(&existingUser).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or phone number already exists."})
		return
	}

	// Create user
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user. Please try again."})
		return
	}

	// Upload image
	imageUrl, err := utils.UploadImage(c)
	fmt.Println(input.ImageUrl.Filename)
	if err != nil {
		log.Printf("Error uploading image for user %s: %v", user.ContactEmail, err)
		database.DB.Delete(&user) // Rollback user creation
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image."})
		return
	}
	user.ImageUrl = imageUrl // Assign image URL to user

	// Upload resume
	resume, err := utils.UploadResume(c)
	if err != nil {
		log.Printf("Error uploading resume for user %s: %v", user.ContactEmail, err)
		database.DB.Delete(&user) // Rollback user creation
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload resume."})
		return
	}
	user.Resume = resume // Assign resume to user

	// Save user with image and resume
	if err := database.DB.Save(&user).Error; err != nil {
		// Rollback and cleanup files
		_ = os.Remove(imageUrl)
		_ = os.Remove(resume)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user data."})
		return
	}

	// Success response
	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "User created successfully",
	})
}

func RegisterCompany(c *gin.Context) {
	var input payloads.RegisterCompany

	// Bind form data (handles both JSON and file data)
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error parsing data. Details: %v", err.Error()),
		})
		return
	}

	// Hash password
	hashPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", input.ContactEmail, err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to process password",
		})
		return
	}
	input.Password = hashPassword

	// Create company instance
	company := models.Company{
		ContactInfo: models.ContactInfo{
			Name:         input.Name,
			ContactEmail: input.ContactEmail,
			ContactPhone: input.ContactPhone,
			Address:      input.Address,
		},
		Description: input.Description,
		Password:    input.Password,
	}

	// Attempt to create company in DB
	if err := database.DB.Create(&company).Error; err != nil {
		log.Printf("Unable to create company: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Unable to create company. Details: %v", err.Error()),
		})
		return
	}

	// Upload the image
	imageUrl, err := utils.UploadImage(c)
	if err != nil {
		log.Printf("Error uploading image for company %s: %v", input.Name, err)

		// Rollback by deleting created company if image upload fails
		database.DB.Delete(&company)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error uploading image. Error: %v", err.Error()),
		})
		return
	}

	// Update company with image URL
	company.ImageUrl = imageUrl
	if err := database.DB.Save(&company).Error; err != nil {
		log.Printf("Failed to update user image for %s: %v", input.Name, err)

		// Delete uploaded image file in case of failure
		os.Remove(imageUrl)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user image",
		})
		return
	}

	// Successfully created company response
	c.JSON(http.StatusOK, gin.H{
		"message": "Company created successfully",
		"company": company,
	})
}

func Login(c *gin.Context) {
	type LoginData struct {
		ContactEmail string `form:"contact_email" json:"contact_email"`
		Password     string `form:"password" json:"password"`
	}
	var loginData LoginData

	// Bind JSON to struct
	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Error parsing data. Details: %v", err.Error()),
		})
		return
	}
	fmt.Println(loginData)
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
		"id":      user.ID,
		"name":    user.Name,
		"image":   user.ImageUrl,
		"resume":  user.Resume,
		"email":   user.ContactEmail,
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

func CheckAuth(c *gin.Context) {
	// Get the access_token cookie
	accessToken, err := c.Cookie("access_token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"isAuthenticated": false,
			"message":         "Access token not found",
		})
		return
	}

	// Validate the JWT token
	claims, err := utils.ValidateJWT(accessToken, true)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"isAuthenticated": false,
			"message":         "Invalid or expired token",
		})
		return
	}

	// If the token is valid, the user is authenticated
	c.JSON(http.StatusOK, gin.H{
		"isAuthenticated": true,
		"message":         "User is authenticated",
		"user": gin.H{ // Include user details from claims (if needed)
			"id":    claims.UserID,
			"email": claims.Email,
			// "role":  claims.Role,
		},
	})
}
