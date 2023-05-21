package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, data any) {
	Respond(c, http.StatusOK, data)
}

func Created(c *gin.Context, data any) {
	Respond(c, http.StatusCreated, data)
}

func Respond(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, gin.H{
		"data": data,
	})
}
