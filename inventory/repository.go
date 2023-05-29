package inventory

import "github.com/dlcrush/casa-hub/adapters"

type InventoryRepository struct {
	adapters.MongoRepository[InventoryItem]
}
