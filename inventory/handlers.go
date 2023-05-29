package inventory

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

var repo InventoryRepository

func ListInventoryItemHandler(c *gin.Context) {
	connectDB()

	items, err := repo.All()
	if err != nil {
		panic(err)
	}

	common.OK(c, gin.H{
		"items": items,
	})
}

func GetInventoryItemHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	item, err := repo.Get(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			common.NotFoundError(c, errors.New("item for id not found"))
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"item": *item,
	})
}

func CreateInventoryItemHandler(c *gin.Context) {
	connectDB()

	var item InventoryItem
	err := c.Bind(&item)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	validate := validator.New()
	err = validate.Struct(item)
	if err != nil {
		common.BadRequestError(c, err)
		return
	}

	item.CreatedAt = time.Now().UTC()
	item.CreatedBy = "TODO"

	resp, err := repo.Create(item)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	fmt.Printf("resp.InsertedID %v", resp.InsertedID)

	item.ID = resp.InsertedID.(primitive.ObjectID)

	common.Created(c, gin.H{
		"item": item,
	})
}

func UpdateInventoryItemHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	var item InventoryItem
	err = c.Bind(&item)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	validate := validator.New()
	err = validate.Struct(item)
	if err != nil {
		common.BadRequestError(c, err)
		return
	}

	existing, err := repo.Get(id)
	if err != nil {
		common.BadRequestError(c, errors.New("id not found"))
		return
	}

	item.ID = existing.ID

	now := time.Now().UTC()
	updateUser := "TODO"
	item.CreatedAt = existing.CreatedAt
	item.CreatedBy = existing.CreatedBy
	item.UpdatedAt = &now
	item.UpdatedBy = &updateUser

	_, err = repo.Update(id, item)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"item": item,
	})
}

func connectDB() {
	repo.Collection = common.DB.Collection("inventory_items")
}
