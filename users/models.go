package users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	FirstName string             `json:"firstName" bson:"first_name" validate:"required"`
	LastName  string             `json:"lastName" bson:"last_name" validate:"required"`
	Password  string             `json:"password" bson:"password"`
}

type UserResponse struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username" bson:"username" validate:"required"`
	Email     string             `json:"email" bson:"email" validate:"required,email"`
	FirstName string             `json:"firstName" bson:"first_name" validate:"required"`
	LastName  string             `json:"lastName" bson:"last_name" validate:"required"`
	Password  string             `json:"-" bson:"password"`
}
