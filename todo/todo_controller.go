package todo

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-todo/models"
	"go-todo/todo/dto"
	"go-todo/utils"
	"time"
)

func CreateTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var req dto.CreateToDoInput
		defer cancel()

		//validate the request body
		if err := c.BindJSON(&req); err != nil {
			utils.BadRequestResponse(c, "Binding Error", err.Error())
			return
		}

		//use the validator library to validate required fields
		if validationErr := validate.Struct(&req); validationErr != nil {
			utils.BadRequestResponse(c, "Struct Error", validationErr.Error())
			return
		}

		newTodo := models.ToDo{
			Name:      req.Name,
			IsDone:    false,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserId:    req.UserId,
		}

		result, err := Create(ctx, newTodo)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}

		utils.OkResponse(c, "Success", result)
	}

}

func GetATodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		todoId := c.Param("todoId")
		defer cancel()

		todo, err := GetById(ctx, todoId)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}
		utils.OkResponse(c, "Success", todo)
	}

}

func GetATodoByUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")
		defer cancel()
		todo, err := GetByUser(ctx, userId)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}
		utils.OkResponse(c, "Success", todo)
	}

}

func GetAllTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		todo, err := GetAll(ctx)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}

		utils.OkResponse(c, "Success", todo)
	}

}

func EditTodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		todoId := c.Param("todoId")
		var req dto.UpdateToDoInput
		defer cancel()

		if err := c.BindJSON(&req); err != nil {
			utils.BadRequestResponse(c, "Error", err.Error())
			return
		}
		req.UpdatedAt = time.Now()

		updatedTodo, err := Update(ctx, todoId, req)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}

		utils.OkResponse(c, "Success", updatedTodo)
	}
}

func DeleteATodo() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		todoId := c.Param("todoId")
		defer cancel()

		deletedTodo, err := Delete(ctx, todoId)
		if err != nil {
			utils.ServerErrorResponse(c, "Error", err.Error())
			return
		}
		utils.OkResponse(c, "Delete Success", deletedTodo)
	}
}
