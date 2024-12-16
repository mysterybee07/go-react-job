package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/mysterybee07/go-react-job/database"
	"github.com/mysterybee07/go-react-job/models"
	"github.com/mysterybee07/go-react-job/utils"
)

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to parse JSON body. Details: %s", err.Error())})
		return
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password for user %s: %v", user.ContactEmail, err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to process password"})
		return
	}

	user.Password = hashPassword

	var existingUser models.User
	if err := database.DB.Where("contact_email = ? OR contact_phone = ?", user.ContactEmail, user.ContactPhone).First(&existingUser).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or phone number already exists."})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user. Please try again."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":    user,
		"message": "User created successfully",
	})
}

func Login(c *gin.Context) {
	type LoginData struct {
		ContactEmail string `json:"contact_email"`
		Password     string `json:"password"`
	}
	var loginData LoginData

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Error parsing data. Details: %v", err.Error())})
		return
	}

	var user models.User
	if err := database.DB.Where("contact_email=?", loginData.ContactEmail).First(&user).Error; err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email not found."})
		return
	}
	match := utils.CheckPasswordHash(user.Password, loginData.Password)
	fmt.Println(match)

	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password."})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Login successful",
		"loginData": loginData,
	})
}
