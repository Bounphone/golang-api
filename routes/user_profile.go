package routes

import (
	"golang_rest_api/controllers"
	"github.com/gin-gonic/gin"
)

func UserProfileRoutes(r *gin.Engine){
	r.POST("/userProfile", controllers.CreateUser)
	r.GET("/userProfile/:userId", controllers.GetUserByID)
}