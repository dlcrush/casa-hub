package properties

import (
	"errors"
	"fmt"
	"time"

	"github.com/dlcrush/casa-hub/common"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var repo PropertyRepository

func ListPropertyHandler(c *gin.Context) {
	connectDB()

	properties, err := repo.All()
	if err != nil {
		panic(err)
	}

	common.OK(c, gin.H{
		"properties": properties,
	})
}

func GetPropertyHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	property, err := repo.Get(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			common.NotFoundError(c, errors.New("item for id not found"))
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"property": *property,
	})
}

func CreatePropertyHandler(c *gin.Context) {
	connectDB()

	var property Property
	err := c.Bind(&property)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	validate := validator.New()
	err = validate.Struct(property)
	if err != nil {
		common.BadRequestError(c, err)
		return
	}

	property.CreatedAt = time.Now().UTC()
	property.CreatedBy = "TODO"

	resp, err := repo.Create(property)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	fmt.Printf("resp.InsertedID %v", resp.InsertedID)

	property.ID = resp.InsertedID.(primitive.ObjectID)

	common.Created(c, gin.H{
		"property": property,
	})
}

func UpdatePropertyHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	var property Property
	err = c.Bind(&property)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	validate := validator.New()
	err = validate.Struct(property)
	if err != nil {
		common.BadRequestError(c, err)
		return
	}

	property.ID = id

	now := time.Now().UTC()
	updateUser := "TODO"
	property.UpdatedAt = &now
	property.UpdatedBy = &updateUser

	_, err = repo.Update(id, property)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"property": property,
	})
}

func connectDB() {
	repo.Collection = common.DB.Collection("properties")
}
