package inventory

import (
	"github.com/dlcrush/casa-hub/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryItem struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Property    primitive.ObjectID `json:"property" bson:"property"`
	Name        string             `json:"name" bson:"name" validate:"required"`
	Description string             `json:"description" bson:"description"`
	Notes       string             `json:"notes" bson:"notes"`
	Room        common.RoomEntity  `json:"room" bson:"room"`
	Images      []common.Image     `json:"images" bson:"images"`
	common.Timestamps
	common.Audit
}
