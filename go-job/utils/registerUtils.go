package utils

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/mysterybee07/go-react-job/database"
	"gorm.io/gorm"
)

func RegisterEntity(c *gin.Context, input interface{}, entity interface{}, uniqueFields map[string]interface{}) {
	// Parse JSON body
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf(
			"Failed to parse JSON body. Details: %s", err.Error(),
		)})
		return
	}

	// Extract password for hashing (assumes input has Password field)
	password := reflect.ValueOf(input).Elem().FieldByName("Password").String()
	hashPassword, err := HashPassword(password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to process password",
		})
		return
	}

	// Update password in input
	reflect.ValueOf(input).Elem().FieldByName("Password").SetString(hashPassword)

	// Populate entity with input data
	copier.Copy(entity, input) // Uses "github.com/jinzhu/copier" to map data easily

	// Check for uniqueness based on provided fields
	var existingEntity interface{}
	query := database.DB
	for field, value := range uniqueFields {
		query = query.Or(fmt.Sprintf("%s = ?", field), value)
	}
	if err := query.First(&existingEntity).Error; err != gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email or phone number already exists.",
		})
		return
	}

	// Create entity in database
	if err := database.DB.Create(entity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create entity. Please try again.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Entity created successfully",
		"entity":  entity,
	})
}
