package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mysterybee07/go-react-job/database"
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
