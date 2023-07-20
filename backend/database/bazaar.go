package database

import (
	"context"

	"skyzar-backend/constants"
	"skyzar-backend/structs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateHypixelReadableNames(data structs.HypixelReadableItemNames) error {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("misc")

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

func GetHypixelReadableNames() (structs.HypixelReadableItemNames, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("misc")

	var names structs.HypixelReadableItemNames

	err := collection.FindOne(
		context.TODO(),
		bson.M{"_id": "hypixelReadableItemNames"},
	).Decode(&names)

	return names, err
}

func GetBazaarItem(id string) (structs.SkyblockBazaarItem, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("products")

	var item structs.SkyblockBazaarItem

	regex := bson.M{"$regex": primitive.Regex{Pattern: id, Options: "i"}}
	filter := bson.M{
		"$or": []interface{}{
			bson.M{"_id": regex},
			bson.M{"hypixel_product_id": regex},
		},
	}
	err := collection.FindOne(
		context.TODO(),
		filter,
	).Decode(&item)

	return item, err
}

func UpdateBazaarItems(models []mongo.WriteModel) error {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("products")

	_, err := collection.BulkWrite(
		context.TODO(),
		models,
		options.BulkWrite().SetOrdered(true),
	)
	return err
}