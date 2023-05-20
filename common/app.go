package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() (*gin.Engine, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading dotenv variables %s\n", err.Error())
		return nil, err
	}

	r := gin.Default()

	return r, nil
}
