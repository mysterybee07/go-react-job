package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/controllers"
)

func JobRoutes(router *gin.Engine) {
	jobGroup := router.Group("/jobs")
	{
		jobGroup.POST("", controllers.CreateJob)
		jobGroup.GET("", controllers.GetJobs)
		jobGroup.GET("/:id", controllers.GetJobByID)
		// jobGroup.PUT("/:id", controllers.UpdateJob)
		jobGroup.DELETE("/:id", controllers.DeleteJob)
	}

	users := router.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
	}
}
