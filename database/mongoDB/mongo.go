package mongoDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

func ConnectDatabase() (*mongo.Collection, error) {
	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-1",
		Username:      os.Getenv("MONGODB_USERNAME"),
		Password:      os.Getenv("MONGODB_PASSWORD"),
	}
	clientOptions := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGODB_HOST") + ":" + os.Getenv("MONGODB_PORT") + "/").SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	collection := client.Database("tasker").Collection("tasks")

	log.Println("MONGODB CONNECTED!")
	return collection, nil
}
