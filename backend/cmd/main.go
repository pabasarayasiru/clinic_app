package main

import (
	"clinic_app/config"
	"clinic_app/models"
	"clinic_app/routes"
	"fmt"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := config.ConnectDB()

	db.AutoMigrate(&models.Doctor{})

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	routes.SetupRoutes(r, db)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	host := os.Getenv("SERVER_HOST")
	if host == "" {
		host = "localhost"
	}

	fmt.Printf("Server is running on http://%s:%s\n", host, port)

	r.Run(":" + port)
}
