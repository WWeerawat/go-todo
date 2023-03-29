package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type jsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func OkResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusOK, jsonResponse{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func BadRequestResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusBadRequest, jsonResponse{
		Status:  http.StatusBadRequest,
		Message: message,
		Data:    data,
	})
}

func ServerErrorResponse(c *gin.Context, message string, data any) {
	c.JSON(http.StatusInternalServerError, jsonResponse{
		Status:  http.StatusInternalServerError,
		Message: message,
		Data:    data,
	})
}
