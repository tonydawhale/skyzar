package database

import (
	"context"

	"skyzar-backend/constants"
	"skyzar-backend/logging"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func StartClient() {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(constants.MongoURI),
	)
	if err != nil {
		logging.LogFatal("Failed to initialize mongo client")
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		logging.LogFatal("Failed to ping mongo client, error: " + err.Error())
	}
	logging.Log("Successfully initialized and pinged mongo client")
	MongoClient = client
}