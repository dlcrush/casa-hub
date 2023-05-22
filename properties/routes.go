package properties

import (
	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
)

func PropertiesRoutes(app *gin.Engine) {
	propertyRoutes := []common.Route{
		{
			Name:     "List Properties",
			Method:   "GET",
			URI:      "/property",
			Handlers: gin.HandlersChain{ListPropertyHandler},
		},
		{
			Name:     "Create Properties",
			Method:   "POST",
			URI:      "/property",
			Handlers: gin.HandlersChain{CreatePropertyHandler},
		},
		{
			Name:     "Get Property",
			Method:   "GET",
			URI:      "/property/:id",
			Handlers: gin.HandlersChain{GetPropertyHandler},
		},
		{
			Name:     "Update Property",
			Method:   "PUT",
			URI:      "/property/:id",
			Handlers: gin.HandlersChain{UpdatePropertyHandler},
		},
		// {
		// 	Name:     "Delete Property",
		// 	Method:   "DELETE",
		// 	URI:      "/property/:id",
		// 	Handlers: gin.HandlersChain{},
		// },
	}

	common.AddRoutes(app, propertyRoutes)
}
