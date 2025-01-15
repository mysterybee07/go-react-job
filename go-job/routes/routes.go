package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/controllers"
	"github.com/mysterybee07/go-react-job/middleware"
)

func JobRoutes(router *gin.Engine) {
	jobGroup := router.Group("/jobs")
	{
		jobGroup.POST("", controllers.CreateJob, middleware.AuthMiddleware())
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
		users.GET("/check-auth", controllers.CheckAuth)
		users.POST("/refresh-token", controllers.RefreshToken)
	}

	company := router.Group("/companies")
	{
		company.POST("/register", controllers.RegisterCompany)
		company.GET("/jobs", controllers.GetJobsByCompany)
		company.GET("", controllers.GetAllCompany)
	}
}
