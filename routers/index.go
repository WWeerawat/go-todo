package routers

import (
	"github.com/gin-gonic/gin"
	"go-todo/todo"
	"go-todo/user"
)

func InitRoute() *gin.Engine {
	router := gin.Default()
	//routes
	user.UserRoutes(router)
	todo.TodoRoutes(router)

	return router
}
