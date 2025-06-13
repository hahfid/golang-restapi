package main

import (
	"log"
	"golang-restapi/config"
	"golang-restapi/models"
	"golang-restapi/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env terlebih dahulu
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect ke database
	config.ConnectDatabase()

	// Migrasi DB otomatis (hati-hati di production)
	err = config.DB.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatal("Error running AutoMigrate:", err)
	}

	// Inisialisasi Gin Router
	r := gin.Default()

	// Setup semua route
	routes.SetupRoutes(r)

	// Jalankan server di port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error running server:", err)
	}
}
