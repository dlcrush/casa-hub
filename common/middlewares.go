package common

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func HandleUncaughtError() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		errorMessage := "unknown error"

		if err != nil {
			errorMessage = err.(error).Error()
		}

		HandleError(c, 500, errors.New(errorMessage))
	}
}
