// package main

// import (
// 	"crypto/tls"
// 	"golang_rest_api/configs"
// 	"golang_rest_api/routes"
// 	"net/http"

// 	// "log"
// 	// "os"

// 	"github.com/gin-gonic/gin"
// 	// "github.com/joho/godotenv"
// )

//	func main() {
//		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
//		// err := godotenv.Load("fly.toml")
//		// if err != nil {
//		// 	log.Fatalf("Error loading .env file")
//		// }
//		configs.ConnectDB()
//		r := gin.Default()
//		routes.UserProfileRoutes(r)
//		r.Run()
//		// port := os.Getenv("PORT")
//		// r.Run(":" + port)
//	}
package main

import (
	"golang_rest_api/configs"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	configs.ConnectDB()
	port := os.Getenv("PORT")
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/", helloHandler)
	log.Println("Listing for" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
