package database

import (
	"context"

	"skyzar-backend/constants"
	"skyzar-backend/structs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

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

func GetBazaarItem(id string) (structs.SkyblockBazaarItem, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection(constants.MongoProductCollection)

	var item structs.SkyblockBazaarItem

	filter := bson.M{
		"_id": id,
	}
	err := collection.FindOne(
		context.TODO(),
		filter,
	).Decode(&item)

	return item, err
}

func GetBazaarItemHistory(id string) (structs.SkyblockBazaarItemHistory, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection(constants.MongoProductHistoryCollection)

	var item structs.SkyblockBazaarItemHistory

	filter := bson.M{
		"_id": id,
	}
	err := collection.FindOne(
		context.TODO(),
		filter,
	).Decode(&item)

	return item, err
}

func GetBazaarItemNames() (map[string]string, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection(constants.MongoProductCollection)

	var namesMongo []structs.SkyblockBazaarItemNameFromMongo

	names := make(map[string]string)

	filter := bson.M{}
	opts := options.Find().SetProjection(bson.M{"_id": 1, "display_name": 1}).SetSort(bson.M{"display_name": 1})
	cursor, err := collection.Find(
		context.TODO(),
		filter,
		opts,
	)

	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &namesMongo); err != nil {
		return nil, err
	}
	for _, item := range(namesMongo) {
		names[item.Id] = item.DisplayName
	}
	return names, nil
}

func GetTopCategory(category string, quota float64) ([]structs.SkyblockBazaarTopItem, int64, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection(constants.MongoProductCollection)

	var items []structs.SkyblockBazaarTopItem
	filter := bson.M{category: bson.M{"$gte": quota}}
	opts := options.Find().
		SetProjection(
			bson.M{
				"_id": 1,
				"display_name": 1,
				"buy_price": 1,
				"sell_price": 1,
				"buy_volume": 1,
				"sell_volume": 1,
				"margin": 1,
				"margin_percent": 1,
			}).
		SetSort(bson.M{category: -1})
	cursor, err := collection.Find(
		context.TODO(),
		filter,
		opts,
	)
	if err != nil {
		return nil, 0, err
	}
	if err := cursor.All(context.Background(), &items); err != nil {
		return nil, 0, err
	}

	totalItems, err := collection.CountDocuments(context.Background(), bson.M{})
	if err != nil {
		return nil, 0, err
	}
	return items, totalItems, nil
}