package common

import "time"

type Image struct {
	URL    string  `json:"url" bson:"url"`
	Width  float32 `json:"width" bson:"width"`
	Height float32 `json:"height" bson:"height"`
}

type Timestamps struct {
	CreatedAt time.Time  `json:"createdAt" bson:"created_at"`
	UpdatedAt *time.Time `json:"updatedAt" bson:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" bson:"deleted_at"`
}

type Audit struct {
	CreatedBy string  `json:"createdBy" bson:"created_by"`
	UpdatedBy *string `json:"updatedBy" bson:"updated_by"`
	DeletedBy *string `json:"deletedBy" bson:"deleted_by"`
}
