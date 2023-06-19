package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/arthuruan/training-consultancy/configs"
	"github.com/arthuruan/training-consultancy/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

// TODO: create one file for each action
func CreateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	c.Bind(&user)

	// TODO: validate the request body

	// if err := c.Bind(&user); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "error",
	// 		"data":    err.Error(),
	// 	})
	// 	return
	// }

	newUser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Type:     user.Type,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
