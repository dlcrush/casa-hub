package main

import (
	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var app *gin.Engine
var db *mongo.Database

func init() {
	var err error
	app, err = common.InitApp()
	if err != nil {
		panic(err)
	}

	db, err = common.InitDB()
	if err != nil {
		panic(err)
	}

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Message": "Hello, World!",
		})
	})
}

func main() {
	app.Run(":3001")
}
