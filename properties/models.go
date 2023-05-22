package properties

import (
	"github.com/dlcrush/casa-hub/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Property struct {
	ID          primitive.ObjectID   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string               `json:"name" bson:"name" validate:"required"`
	Owners      []primitive.ObjectID `json:"owners" bson:"owners" validate:"required"`
	Address     string               `json:"address" bson:"address" validate:"required"`
	Description string               `json:"description" bson:"description" validate:"required"`
	Type        string               `json:"type" bson:"type" validate:"required"`
	Images      []common.Image       `json:"images" bson:"images"`
	Notes       string               `json:"notes" bson:"notes"`
	Rooms       []Room               `json:"rooms" bson:"rooms"`
	common.Timestamps
	common.Audit
}

type Room struct {
	Key         string         `json:"key" bson:"key"`
	Name        string         `json:"name" bson:"name"`
	Description string         `json:"description" bson:"description"`
	Notes       string         `json:"notes" bson:"notes"`
	Images      []common.Image `json:"images" bson:"images"`
	common.Timestamps
	common.Audit
}
