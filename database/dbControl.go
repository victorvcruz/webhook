package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Task struct {
	ID     float64 `bson:"_id"`
	Thread string  `bson:"thread"`
}

func Insert(id float64, thread string, collection *mongo.Collection) error {
	task := &Task{
		ID:     id,
		Thread: thread,
	}

	_, err := collection.InsertOne(context.TODO(), task)
	return err
}

func Find(id float64, collection *mongo.Collection) (map[string]interface{}, error) {

	var result bson.M
	err := collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}

	var mapJson = make(map[string]interface{})
	for key, value := range result {
		mapJson[key] = value
	}

	return mapJson, nil
}
