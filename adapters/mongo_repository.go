package adapters

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository[T any] struct {
	Collection *mongo.Collection
}

type IMongoRepository[T any] interface {
	All() (*[]T, error)
	Find(filter bson.D) (*[]T, error)
	FindOne(filter bson.D) (*T, error)
	Get(id primitive.ObjectID) (*T, error)
	Create(item T) (*mongo.InsertOneResult, error)
	Update(id primitive.ObjectID, item T) (*mongo.UpdateResult, error)
	Delete(id primitive.ObjectID) (*mongo.DeleteResult, error)
}

func (r MongoRepository[T]) All() (*[]T, error) {
	return r.Find(bson.D{})
}

func (r MongoRepository[T]) Find(filter bson.D) (*[]T, error) {
	cursor, err := r.Collection.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Error doing Mongo Find %s\n", err.Error())
		return nil, err
	}

	var items []T
	if cursor.RemainingBatchLength() < 1 {
		items = make([]T, 0)
		return &items, nil
	}

	err = cursor.All(context.TODO(), &items)
	if err != nil {
		fmt.Printf("Error on cursor.All %s\n", err.Error())
		return nil, err
	}

	return &items, nil
}

func (r MongoRepository[T]) FindOne(filter bson.D) (*T, error) {
	resp := r.Collection.FindOne(context.TODO(), filter)

	var item T
	err := resp.Decode(&item)
	if err != nil {
		fmt.Printf("Error decoding result %s\n", err.Error())
		return nil, err
	}

	return &item, nil
}

func (r MongoRepository[T]) Get(id primitive.ObjectID) (*T, error) {
	return r.FindOne(bson.D{{Key: "_id", Value: id}})
}

func (r MongoRepository[T]) Create(item T) (*mongo.InsertOneResult, error) {
	return r.Collection.InsertOne(context.TODO(), item)
}

func (r MongoRepository[T]) Update(id primitive.ObjectID, item T) (*mongo.UpdateResult, error) {
	return r.Collection.ReplaceOne(context.TODO(), bson.D{{Key: "_id", Value: id}}, item)
}

func (r MongoRepository[T]) Delete(id primitive.ObjectID) (*mongo.DeleteResult, error) {
	return r.Collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: id}})
}
