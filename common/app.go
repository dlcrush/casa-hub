package common

import (
	"github.com/gin-gonic/gin"
)

func InitApp() (*gin.Engine, error) {
	r := gin.Default()

	r.SetTrustedProxies(nil)

	r.Use(gin.CustomRecovery(HandleUncaughtError()))

	return r, nil
}
