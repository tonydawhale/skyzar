package database

import (
	"context"

	"skyzar-database/constants"
	"skyzar-database/structs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateHypixelReadableNames(data structs.HypixelReadableItemNames) error {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection(constants.MongoReadableNamesCollection)

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": "hypixelReadableItemNames"},
		bson.M{
			"$set": data,
		},
		options.Update().SetUpsert(true),
	)
	return err
}

func BulkWriteUpdate(models []mongo.WriteModel, col string) error {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection(col)

	_, err := collection.BulkWrite(
		context.TODO(),
		models,
		options.BulkWrite().SetOrdered(true),
	)
	return err
}

func GetHypixelReadableNames() (structs.HypixelReadableItemNames, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection(constants.MongoReadableNamesCollection)

	var names structs.HypixelReadableItemNames

	err := collection.FindOne(
		context.TODO(),
		bson.M{"_id": "hypixelReadableItemNames"},
	).Decode(&names)

	return names, err
}