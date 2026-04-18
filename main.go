package main

import (
	"log"

	"belajar-go/config"
	"belajar-go/providers"
	"belajar-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	r := gin.Default()

	// koneksi DB dulu
	config.ConnectDB()
	config.RunMigration()
	db := config.DB

	// dependency injection
	container := providers.NewContainer(db)

	// kirim controller ke routes
	routes.UserRoutes(r, container.UserController)

	r.Run(":8080")
}
