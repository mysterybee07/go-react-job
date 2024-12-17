package utils

import (
	"fmt"
	"strconv"

	"github.com/mysterybee07/go-react-job/database"
	"github.com/mysterybee07/go-react-job/models"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash checks if the provided password matches the hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// CompanyLogin handles login logic for a company
func CompanyLogin(email, password string) (token, refreshToken string, err error) {
	// Fetch company from database
	var company models.Company
	if err := database.DB.Where("contact_email = ?", email).First(&company).Error; err != nil {
		return "", "", err // Return error if company not found
	}

	// Validate password
	match := CheckPasswordHash(password, company.Password)
	if !match {
		return "", "", fmt.Errorf("invalid password")
	}

	// Generate JWT tokens
	companyID := strconv.Itoa(int(company.ID))
	token, refreshToken, err = GenerateJWT(companyID, company.Name, company.ContactEmail)
	if err != nil {
		return "", "", fmt.Errorf("token generation failed: %v", err)
	}

	return token, refreshToken, nil
}
