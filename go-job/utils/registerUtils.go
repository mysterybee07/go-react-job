package utils

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func UploadFile(c *gin.Context, formFieldName string, validMimeTypes []string, uploadDir string) (string, error) {
	// Retrieve the uploaded file
	file, err := c.FormFile(formFieldName)
	if err != nil {
		log.Println("Failed to get form file:", err)
		return "", fmt.Errorf("failed to get form file: %v", err)
	}

	// Validate the file's MIME type
	fileType := file.Header.Get("Content-Type")
	isValidMime := false
	for _, validType := range validMimeTypes {
		if fileType == validType {
			isValidMime = true
			break
		}
	}
	if !isValidMime {
		log.Printf("Uploaded file has an invalid MIME type: %s", fileType)
		return "", fmt.Errorf("invalid file type: %s", fileType)
	}

	// Generate a random file name and sanitize it
	fileName := RandLetter(5) + "-" + SanitizeFileName(file.Filename)
	filePath := filepath.Join(uploadDir, fileName)

	// Ensure the upload directory exists
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Println("Failed to create upload directory:", err)
		return "", fmt.Errorf("failed to create upload directory: %v", err)
	}

	// Save the uploaded file
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Failed to save uploaded file:", err)
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return filePath, nil
}

func UploadImage(c *gin.Context) (string, error) {
	validMimeTypes := []string{"image/jpeg", "image/png", "image/gif"}
	return UploadFile(c, "image_url", validMimeTypes, "./uploads/images")
}

func UploadResume(c *gin.Context) (string, error) {
	validMimeTypes := []string{"application/pdf", "application/msword", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}
	return UploadFile(c, "resume", validMimeTypes, "./uploads/resumes")
}

// RandLetter generates a random string of length n using letters.
func RandLetter(n int) string {
	b := make([]rune, n)
	rand.Seed(time.Now().UnixNano()) // time.Now().UnixNano() is int64, suitable for math/rand
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func SanitizeFileName(fileName string) string {
	// Remove any characters that are not alphanumeric, dot, or underscore
	re := regexp.MustCompile(`[^a-zA-Z0-9._-]`)
	return strings.ToLower(re.ReplaceAllString(fileName, "_"))
}

// func RegisterEntity(c *gin.Context, input interface{}, entity interface{}, uniqueFields map[string]interface{}) {
// 	// Parse JSON body
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(
// 			"Failed to parse JSON body. Details: %s", err.Error(),
// 		)})
// 		return
// 	}

// 	// Extract password for hashing (assumes input has Password field)
// 	password := reflect.ValueOf(input).Elem().FieldByName("Password").String()
// 	hashPassword, err := HashPassword(password)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "Failed to process password",
// 		})
// 		return
// 	}

// 	// Update password in input
// 	reflect.ValueOf(input).Elem().FieldByName("Password").SetString(hashPassword)

// 	// Populate entity with input data
// 	copier.Copy(entity, input) // Uses "github.com/jinzhu/copier" to map data easily

// 	// Check for uniqueness based on provided fields
// 	var existingEntity interface{}
// 	query := database.DB
// 	for field, value := range uniqueFields {
// 		query = query.Or(fmt.Sprintf("%s = ?", field), value)
// 	}
// 	if err := query.First(&existingEntity).Error; err != gorm.ErrRecordNotFound {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Email or phone number already exists.",
// 		})
// 		return
// 	}

// 	// Create entity in database
// 	if err := database.DB.Create(entity).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Could not create entity. Please try again.",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Entity created successfully",
// 		"entity":  entity,
// 	})
// }
