package user

import "github.com/gin-gonic/gin"

func UserRoutes(router *gin.Engine) {
	router.POST("/user", CreateUser())
	router.GET("/users", GetAllUser())
	router.GET("/user/:userId", GetAUser())
	router.PATCH("/user/:userId", EditUser())
	router.DELETE("/user/:userId", DeleteAUser())
}
