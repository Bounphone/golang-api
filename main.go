package main

import (
	// "crypto/tls"
	"golang_rest_api/configs"
	"golang_rest_api/routes"
	"net/http"

	// "golang_rest_api/configs"
	// "golang_rest_api/routes"
	// "net/http"

	// "log"
	// "os"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

	func main() {
		// http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
		// err := godotenv.Load("fly.toml")
		// if err != nil {
		// 	log.Fatalf("Error loading .env file")
		// }
		configs.ConnectDB()
		r := gin.Default()
		// r.SetTrustedProxies([]string{"192.168.1.2"})
		routes.UserProfileRoutes(r)
		r.Run()
		// port := os.Getenv("PORT")
		// r.Run(":" + port)
	}