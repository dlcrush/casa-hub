package tasks

import (
	"time"

	"github.com/dlcrush/casa-hub/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty"`
	Property    primitive.ObjectID  `json:"property" bson:"property" validate:"required"`
	Assignees   []common.UserEntity `json:"assignees" bson:"assignees" validate:"required"`
	Name        string              `json:"name" bson:"name" validate:"required"`
	Description string              `json:"description" bson:"description"`
	Date        time.Time           `json:"date" bson:"date"`
	DueDate     *time.Time          `json:"dueDate" bson:"dueDate:"`
	Notes       string              `json:"notes" bson:"notes"`
	common.Timestamps
	common.Audit
}
