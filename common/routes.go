package common

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name     string
	Method   string
	URI      string
	Handlers gin.HandlersChain
}

func AddRoute(app *gin.Engine, route Route) {
	switch strings.ToLower(route.Method) {
	case "get":
		app.GET(route.URI, route.Handlers...)
	case "post":
		app.POST(route.URI, route.Handlers...)
	case "put":
		app.PUT(route.URI, route.Handlers...)
	case "patch":
		app.PATCH(route.URI, route.Handlers...)
	case "delete":
		app.DELETE(route.URI, route.Handlers...)
	case "options":
		app.OPTIONS(route.Method, route.Handlers...)
	}
}

func AddRoutes(app *gin.Engine, routes []Route) {
	for _, r := range routes {
		AddRoute(app, r)
	}
}
