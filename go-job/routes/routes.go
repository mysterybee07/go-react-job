package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/controllers"
	"github.com/mysterybee07/go-react-job/middleware"
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
		users.POST("/register", controllers.RegisterUser)
		users.POST("/login", controllers.Login)
		users.POST("/logout", controllers.Logout)
		users.GET("/authorize", middleware.AuthMiddleware(), controllers.AuthorizedUser)
		users.POST("/refresh-token", controllers.RefreshToken)
	}

	company := router.Group("/company")
	{
		company.POST("/register", controllers.RegisterCompany)
	}
}
