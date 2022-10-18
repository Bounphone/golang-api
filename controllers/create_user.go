package controllers

import (
	"context"
	"golang_rest_api/configs"
	"golang_rest_api/models"
	"golang_rest_api/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	//validate the request body
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, responses.UserProfileResponse{Status: http.StatusBadRequest, Message: "error", Data: err.Error()})
		return
	}
	// if err := c.BodyParser(&user); err != nil {
	// 	return c.Status(http.StatusBadRequest).JSON(responses.UserProfileResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	// }

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(http.StatusInternalServerError, responses.UserProfileResponse{Status: http.StatusBadRequest, Message: "error", Data: validationErr.Error()})
		return
		// return c.Status(http.StatusBadRequest).JSON(responses.UserProfileResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newUser := models.User{
		Name:     user.Name,
		Location: user.Location,
		Title:    user.Title,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.UserProfileResponse{Status: http.StatusInternalServerError, Message: "error", Data: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.UserProfileResponse{Status: http.StatusOK, Message: "success", Data: result})
	return
}

func GetUserByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	userId := c.Param("userId")
	var user models.User
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(userId)
	err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		message := err.Error()
		if err.Error() == "mongo: no documents in result" {
			message = "Your document was not found in DB"
		}
		c.JSON(http.StatusInternalServerError, responses.UserProfileResponse{Status: http.StatusInternalServerError, Message: "error", Data: message})
		return
	}
	c.JSON(http.StatusOK, responses.UserProfileResponse{Status: http.StatusOK, Message: "success", Data: user})
		return
}
