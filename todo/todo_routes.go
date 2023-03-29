package todo

import "github.com/gin-gonic/gin"

func TodoRoutes(router *gin.Engine) {
	router.POST("/todo", CreateTodo())
	router.GET("/todos", GetAllTodo())
	router.GET("/todo/:todoId", GetATodo())
	router.GET("/todo/user/:userId", GetATodoByUser())
	router.PATCH("/todo/:todoId", EditTodo())
	router.DELETE("/todo/:todoId", DeleteATodo())
}
