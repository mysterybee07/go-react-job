package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/database"
	"github.com/mysterybee07/go-react-job/models"
	"github.com/mysterybee07/go-react-job/utils"
	"gorm.io/gorm"
)

// Create a new job
func CreateJob(c *gin.Context) {
	var job models.Job
	// Parse the JSON body
	if err := c.ShouldBindJSON(&job); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("Error creating job. Details: %v", err.Error()),
		})
		return
	}
	// Return the created job
	c.JSON(http.StatusOK, gin.H{
		"message": "job created successfully",
		"job":     job,
	})
}

func GetJobs(c *gin.Context) {
	var jobs []models.Job

	// Fetch the 'limit' query parameter
	limitStr := c.Query("limit") // Get 'limit' from the query string
	limit := 0                   // Default limit: 0 (no limit)

	// Parse the 'limit' parameter to an integer
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit parameter"})
			return
		}
		limit = parsedLimit
	}

	// Modify the database query to include the limit
	query := database.DB
	if limit > 0 {
		query = query.Limit(limit)
	}

	if err := query.Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the jobs
	c.JSON(http.StatusOK, gin.H{
		"jobs": jobs,
	})
}

func GetJobByID(c *gin.Context) {
	id := c.Param("id")
	var job models.Job

	// Preload the associated Company details along with Job
	if err := database.DB.Preload("Company").First(&job, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the job details along with the associated company
	c.JSON(http.StatusOK, job)
}

func DeleteJob(c *gin.Context) {
	id := c.Param("id")
	var job models.Job
	if err := database.DB.First(&job, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Job not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Soft delete the job
	if err := database.DB.Delete(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Job deleted successfully"})
}

func GetJobsByCompany(c *gin.Context) {
	// Extract company ID from JWT token
	companyID, err := utils.ExtractUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized: " + err.Error(),
		})
		return
	}

	// Fetch jobs for the company
	var jobs []models.Job
	if err := database.DB.Where("company_id = ?", companyID).Preload("Company").Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch jobs.",
		})
		return
	}

	// Return the jobs as JSON response
	c.JSON(http.StatusOK, gin.H{
		"company_id": companyID,
		"jobs":       jobs,
	})
}
