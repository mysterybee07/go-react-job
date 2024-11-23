package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/controllers"
)

func JobRoutes(router *gin.Engine) {
	jobGroup := router.Group("/jobs")
	{
		jobGroup.POST("", controllers.CreateJob)     // Create a new job
		jobGroup.GET("", controllers.GetJobs)        // Get all jobs
		jobGroup.GET("/:id", controllers.GetJobByID) // Get a single job by ID
		// jobGroup.PUT("/:id", controllers.UpdateJob)    // Update a job by ID
		jobGroup.DELETE("/:id", controllers.DeleteJob) // Delete a job by ID
	}
}
