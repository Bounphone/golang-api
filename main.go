package main

import (
	"golang_rest_api/configs"
	"golang_rest_api/routes"
	// "log"
	// "os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load("fly.toml")
	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }
	configs.ConnectDB()
	r := gin.Default()
	routes.UserProfileRoutes(r)
	r.Run()
	// port := os.Getenv("PORT")
	// r.Run(":" + port)
}
