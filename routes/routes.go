package routes

import (
	"github.com/dlcrush/casa-hub/inventory"
	"github.com/dlcrush/casa-hub/properties"
	"github.com/dlcrush/casa-hub/tasks"
	"github.com/dlcrush/casa-hub/users"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	users.UserRoutes(app)
	properties.PropertiesRoutes(app)
	tasks.TaskRoutes(app)
	inventory.InventoryRoutes(app)
}
