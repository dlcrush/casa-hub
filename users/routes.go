package users

import (
	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
)

func UserRoutes(app *gin.Engine) {
	userRoutes := []common.Route{
		{
			Name:     "List Users",
			Method:   "GET",
			URI:      "/user",
			Handlers: gin.HandlersChain{ListUserHandler},
		},
		{
			Name:     "Create User",
			Method:   "POST",
			URI:      "/user",
			Handlers: gin.HandlersChain{CreateUserHandler},
		},
		{
			Name:     "Get User",
			Method:   "GET",
			URI:      "/user/:id",
			Handlers: gin.HandlersChain{GetUserHandler},
		},
		{
			Name:     "Update User",
			Method:   "PUT",
			URI:      "/user/:id",
			Handlers: gin.HandlersChain{UpdateUserHandler},
		},
		{
			Name:     "User Login",
			Method:   "POST",
			URI:      "/login",
			Handlers: gin.HandlersChain{LoginUserHandler},
		},
	}

	common.AddRoutes(app, userRoutes)
}
