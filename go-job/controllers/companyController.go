package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/database"
	"github.com/mysterybee07/go-react-job/models"
)

func GetAllCompany(c *gin.Context) {
	var companies []models.Company

	if err := database.DB.Find(&companies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("No Company Found. Details: %v", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"companies": companies,
	})
}
