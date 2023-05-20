package common

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() (*mongo.Database, error) {
	opts := getConnectionOptions()

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		fmt.Printf("Error connecting to Mongo DB %s\n", err.Error())
		return nil, err
	}

	return client.Database("casaHubDB"), nil
}

func getConnectionOptions() *options.ClientOptions {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.jociec8.mongodb.net/?retryWrites=true&w=majority", os.Getenv("MONGO_DB_USERNAME"), os.Getenv("MONGO_DB_PASSWORD"))

	return options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
}
