package users

import (
	"errors"

	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var repo UserRepository

func ListUserHandler(c *gin.Context) {
	connectDB()

	users, err := repo.All()
	if err != nil {
		panic(err)
	}

	common.OK(c, gin.H{
		"users": toUsersWithoutPassword(*users),
	})
}

func CreateUserHandler(c *gin.Context) {
	connectDB()

	var user User
	err := c.Bind(&user)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	resp, err := repo.Create(user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	user.ID = resp.InsertedID.(primitive.ObjectID)

	common.Created(c, gin.H{
		"user": toUserWithoutPassword(user),
	})
}

func GetUserHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	user, err := repo.Get(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			common.NotFoundError(c, errors.New("item for id not found"))
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"user": toUserWithoutPassword(*user),
	})
}

func UpdateUserHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	var user User
	err = c.Bind(&user)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	// Make sure user isn't trying to write a new ID
	user.ID = id

	_, err = repo.Update(id, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"user": toUserWithoutPassword(user),
	})
}

func toUsersWithoutPassword(users []User) []UserResponse {
	var userResponses []UserResponse

	for _, usr := range users {
		userResponses = append(userResponses, toUserWithoutPassword(usr))
	}

	return userResponses
}

func toUserWithoutPassword(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func connectDB() {
	repo.Collection = common.DB.Collection("users")
}
