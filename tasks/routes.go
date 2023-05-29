package tasks

import (
	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
)

func TaskRoutes(app *gin.Engine) {
	taskRoutes := []common.Route{
		{
			Name:     "List Tasks",
			Method:   "GET",
			URI:      "/task",
			Handlers: gin.HandlersChain{ListTaskHandler},
		},
		{
			Name:     "Create Task",
			Method:   "POST",
			URI:      "/task",
			Handlers: gin.HandlersChain{CreateTaskHandler},
		},
		{
			Name:     "Get Task",
			Method:   "GET",
			URI:      "/task/:id",
			Handlers: gin.HandlersChain{GetTaskHandler},
		},
		{
			Name:     "Update Task",
			Method:   "PUT",
			URI:      "/task/:id",
			Handlers: gin.HandlersChain{UpdateTaskHandler},
		},
	}

	common.AddRoutes(app, taskRoutes)
}
