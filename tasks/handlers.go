package tasks

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

var repo TaskRepository

func ListTaskHandler(c *gin.Context) {
	connectDB()

	tasks, err := repo.All()
	if err != nil {
		panic(err)
	}

	common.OK(c, gin.H{
		"tasks": tasks,
	})
}

func GetTaskHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	task, err := repo.Get(id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			common.NotFoundError(c, errors.New("item for id not found"))
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"task": *task,
	})
}

func CreateTaskHandler(c *gin.Context) {
	connectDB()

	var task Task
	err := c.Bind(&task)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	validate := validator.New()
	err = validate.Struct(task)
	if err != nil {
		common.BadRequestError(c, err)
		return
	}

	task.CreatedAt = time.Now().UTC()
	task.CreatedBy = "TODO"

	resp, err := repo.Create(task)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	fmt.Printf("resp.InsertedID %v", resp.InsertedID)

	task.ID = resp.InsertedID.(primitive.ObjectID)

	common.Created(c, gin.H{
		"task": task,
	})
}

func UpdateTaskHandler(c *gin.Context) {
	connectDB()

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		common.BadRequestError(c, errors.New("invalid id param"))
		return
	}

	var task Task
	err = c.Bind(&task)
	if err != nil {
		common.BadRequestError(c, errors.New("invalid body"))
		return
	}

	validate := validator.New()
	err = validate.Struct(task)
	if err != nil {
		common.BadRequestError(c, err)
		return
	}

	existing, err := repo.Get(id)
	if err != nil {
		common.BadRequestError(c, errors.New("id not found"))
		return
	}

	task.ID = existing.ID

	now := time.Now().UTC()
	updateUser := "TODO"
	task.CreatedAt = existing.CreatedAt
	task.CreatedBy = existing.CreatedBy
	task.UpdatedAt = &now
	task.UpdatedBy = &updateUser

	_, err = repo.Update(id, task)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			common.BadRequestError(c, err)
			return
		}
		panic(err)
	}

	common.OK(c, gin.H{
		"task": task,
	})
}

func connectDB() {
	repo.Collection = common.DB.Collection("tasks")
}
