package main

import (
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	mongoDB2 "webhooks-chat/database/mongoDB"
	httpserver "webhooks-chat/http-server"
)

var mongoDB *mongo.Collection
var port string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoDB, err = mongoDB2.ConnectDatabase()
	if err != nil {
		log.Fatal("Error connecting database mongo")
	}

	port = "8080"
	env := os.Getenv("API_PORT")
	if env != "" {
		port = env
	}
}

func main() {

	api := httpserver.API{mongoDB}

	api.Run(port)
}
