package inventory

import (
	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
)

func InventoryRoutes(app *gin.Engine) {
	inventoryRoutes := []common.Route{
		{
			Name:     "List Inventory Items",
			Method:   "GET",
			URI:      "/inventory",
			Handlers: gin.HandlersChain{ListInventoryItemHandler},
		},
		{
			Name:     "Create Inventory Item",
			Method:   "POST",
			URI:      "/inventory",
			Handlers: gin.HandlersChain{GetInventoryItemHandler},
		},
		{
			Name:     "Get Inventory Item",
			Method:   "GET",
			URI:      "/inventory/:id",
			Handlers: gin.HandlersChain{CreateInventoryItemHandler},
		},
		{
			Name:     "Update Iventory Item",
			Method:   "PUT",
			URI:      "/inventory/:id",
			Handlers: gin.HandlersChain{UpdateInventoryItemHandler},
		},
	}

	common.AddRoutes(app, inventoryRoutes)
}
