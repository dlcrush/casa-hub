package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequestError(c *gin.Context, err error) {
	HandleError(c, http.StatusBadRequest, err)
}

func InternalServerError(c *gin.Context, err error) {
	HandleError(c, http.StatusInternalServerError, err)
}

func UnauthorizedError(c *gin.Context, err error) {
	HandleError(c, http.StatusUnauthorized, err)
}

func NotFoundError(c *gin.Context, err error) {
	HandleError(c, http.StatusNotFound, err)
}

func HandleError(c *gin.Context, statusCode int, err error) {
	errorMessage := err.Error()
	if len(errorMessage) < 1 {
		errorMessage = "Unknown error occured"
	}
	c.AbortWithStatusJSON(statusCode, gin.H{
		"data": gin.H{
			"error": gin.H{
				"code":    statusCode,
				"message": errorMessage,
			},
		},
	})
}
