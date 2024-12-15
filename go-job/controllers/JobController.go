package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/database"
	"github.com/mysterybee07/go-react-job/models"
	"gorm.io/gorm"
)

// Create a new job
func CreateJob(c *gin.Context) {
	var jobRequest struct {
		Title       string `json:"title"`
		Type        string `json:"type"`
		Location    string `json:"location"`
		Description string `json:"description"`
		Salary      string `json:"salary"`
		Company     struct {
			Name         string `json:"name"`
			Description  string `json:"description"`
			ContactEmail string `json:"contactEmail"`
			ContactPhone string `json:"contactPhone"`
		} `json:"company"`
	}

	// Parse the JSON body
	if err := c.ShouldBindJSON(&jobRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the company already exists
	var company models.Company
	if err := database.DB.Where("name = ?", jobRequest.Company.Name).First(&company).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create the company if it doesn't exist
			company = models.Company{
				Description: jobRequest.Company.Description,
				ContactInfo: models.ContactInfo{
					Name:         jobRequest.Company.Name,
					ContactEmail: jobRequest.Company.ContactEmail,
					ContactPhone: jobRequest.Company.ContactPhone,
				},
			}
			if err := database.DB.Create(&company).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create company"})
				return
			}
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching company"})
			return
		}
	}

	// Create the job and associate the company
	job := models.Job{
		Title:       jobRequest.Title,
		Type:        jobRequest.Type,
		Location:    jobRequest.Location,
		Description: jobRequest.Description,
		Salary:      jobRequest.Salary,
		CompanyID:   company.ID, // Associate the company ID
	}

	if err := database.DB.Create(&job).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create job"})
		return
	}

	// Fetch the job along with the company details
	var createdJob models.Job
	if err := database.DB.Preload("Company").First(&createdJob, job.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch job details"})
		return
	}

	// Return the created job
	c.JSON(http.StatusOK, gin.H{
		"job": createdJob,
	})
}

func GetJobs(c *gin.Context) {
	var jobs []models.Job
	if err := database.DB.Find(&jobs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, jobs)
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
