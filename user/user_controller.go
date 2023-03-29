package user

import (
	"context"
	"go-todo/models"
	"go-todo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&user); err != nil {

			utils.BadRequestResponse(c, "Error", err.Error())
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			utils.BadRequestResponse(c, "Error", validationErr.Error())
			return
		}

		newUser := models.User{
			Email:     user.Email,
			Password:  user.Password,
			Name:      user.Name,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		result, err := Create(ctx, newUser)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}

		utils.OkResponse(c, "Success", result)
	}

}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		user, err := GetById(ctx, userId)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}
		utils.OkResponse(c, "Success", user)
	}

}

func GetAllUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		user, err := GetAll(ctx)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}

		utils.OkResponse(c, "Success", user)
	}

}

func EditUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			utils.BadRequestResponse(c, "Error", err.Error())
			return
		}

		update := bson.M{
			"Email":     user.Email,
			"Password":  user.Password,
			"Name":      user.Name,
			"CreatedAt": user.CreatedAt,
			"UpdatedAt": time.Now(),
		}

		updatedUser, err := Update(ctx, userId, update)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}

		utils.OkResponse(c, "Success", updatedUser)
	}
}

func DeleteAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()

		deletedUser, err := Delete(ctx, userId)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}
		utils.OkResponse(c, "Delete Success", deletedUser)
	}
}
