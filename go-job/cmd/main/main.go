package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/database"
	"github.com/mysterybee07/go-react-job/routes"
	"github.com/mysterybee07/go-react-job/utils"
)

func init() {
	utils.LoadEnv()
	database.ConnectDB()
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.Default()
	// Enable CORS for all routes (adjust options as needed)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Allowed methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		AllowCredentials: true,
	}))

	routes.JobRoutes(r)

	// Example route
	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start the server
	err := r.Run(":" + port) // listen and serve on the specified port
	if err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
