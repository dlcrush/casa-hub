package users

import (
	"errors"

	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	if len(user.Password) < 8 {
		common.BadRequestError(c, errors.New("password length too short"))
		return
	}

	passwordHash, err := hashPassword(user.Password)
	if err != nil {
		common.InternalServerError(c, errors.New("error creating password hash"))
		return
	}

	user.Password = passwordHash

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		common.BadRequestError(c, err)
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

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		common.BadRequestError(c, err)
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

func LoginUserHandler(c *gin.Context) {
	connectDB()

	var body LoginBody
	err := c.Bind(&body)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	validate := validator.New()
	err = validate.Struct(body)
	if err != nil {
		common.BadRequestError(c, err)
		return
	}

	user, err := repo.FindOne(bson.D{{Key: "email", Value: body.Email}})
	if err != nil {
		common.UnauthorizedError(c, err)
		return
	}

	if doPasswordsMatch(user.Password, body.Password) == false {
		common.UnauthorizedError(c, errors.New("Invalid credentials"))
		return
	}

	token, err := CreateJWT(*user)
	if err != nil {
		common.InternalServerError(c, err)
		return
	}

	common.OK(c, gin.H{
		"token": token,
	})
}

func doPasswordsMatch(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	var passwordBytes = []byte(password)

	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
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
