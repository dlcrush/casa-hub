package main

import (
	"fmt"

	"github.com/dlcrush/casa-hub/common"
	"github.com/dlcrush/casa-hub/properties"
	"github.com/dlcrush/casa-hub/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var app *gin.Engine

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading dotenv variables %s\n", err.Error())
		panic(err)
	}

	app, err = common.InitApp()
	if err != nil {
		panic(err)
	}

	common.OpenMongoConnection()

	users.UserRoutes(app)
	properties.PropertiesRoutes(app)
}

func main() {
	app.Run("localhost:3001")
	defer common.CloseMongoConnection()
}
