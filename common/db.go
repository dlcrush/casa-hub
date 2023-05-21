package common

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func InitDB() (*mongo.Database, error) {
	var err error
	Client, err = GetMongoConnection()
	if err != nil {
		return nil, err
	}

	DB = Client.Database("casaHubDB")

	return DB, nil
}

func GetMongoConnection() (*mongo.Client, error) {
	if Client != nil {
		return Client, nil
	}

	var err error
	Client, err = mongo.Connect(context.TODO(), getConnectionOptions())
	if err != nil {
		return nil, err
	}

	return Client, nil
}

func OpenMongoConnection() {
	InitDB()
}

func CloseMongoConnection() {
	fmt.Println("CloseMongoConnection()")
	if Client != nil {
		Client.Disconnect(context.TODO())
		Client = nil
	}
}

func GetCollection(collection string) *mongo.Collection {
	return DB.Collection(collection)
}

func getConnectionOptions() *options.ClientOptions {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.jociec8.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_DB_USERNAME"), os.Getenv("MONGO_DB_PASSWORD"))

	return options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
}
