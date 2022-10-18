package main

import (
	"crypto/tls"
	// "golang_rest_api/configs"
	"golang_rest_api/routes"
	"net/http"

	// "log"
	// "os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

	func main() {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		// err := godotenv.Load("fly.toml")
		// if err != nil {
		// 	log.Fatalf("Error loading .env file")
		// }
		// configs.ConnectDB()
		r := gin.Default()
		// r.SetTrustedProxies([]string{"192.168.1.2"})
		routes.UserProfileRoutes(r)
		r.Run()
		// port := os.Getenv("PORT")
		// r.Run(":" + port)
	}