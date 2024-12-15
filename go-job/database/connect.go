package database

import (
	"fmt"
	"log"
	"os"

	"github.com/mysterybee07/go-react-job/models"
	"github.com/mysterybee07/go-react-job/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	utils.LoadEnv()

	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Validate required environment variables
	if dbUser == "" || dbHost == "" || dbName == "" {
		log.Fatal("One or more required environment variables (DB_USER, DB_HOST, DB_NAME) are not set")
	}

	// Handle case where password is not set (leave it empty)
	if dbPassword == "" {
		dbPassword = ""
	}

	// Build the DSN (Data Source Name) dynamically
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbName)

	// Open a connection to the database using GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Log successful connection
	log.Println("Connected to database successfully")

	if err := DB.AutoMigrate(
		models.User{},
		models.Company{},
		models.Job{},
	); err != nil {
		panic("Error migrating to database")
	}

}
