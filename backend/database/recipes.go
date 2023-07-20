package database

import (
	"context"

	"skyzar-backend/constants"
	"skyzar-backend/structs"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRecipes() ([]structs.SkyblockItemRecipe, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("recipes")

	var recipes []structs.SkyblockItemRecipe

	filter := bson.M{}
	opts := options.Find().SetProjection(bson.M{"_id": 0})
	cursor, err := collection.Find(
		context.TODO(),
		filter,
		opts,
	)
	if err != nil {
		return recipes, err
	}
	if err := cursor.All(context.Background(), &recipes); err != nil {
		return recipes, err
	}

	return recipes, nil
}

func CreateRecipe(data structs.SkyblockItemRecipe) error {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("recipes")

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"itemId": data.ItemID},
		bson.M{
			"$setOnInsert": bson.M{
				"_id": primitive.NewObjectID(),
			},
			"$set": data,
		},
		options.Update().SetUpsert(true),
	)
	return err
}

func GetRecipe(id string) (structs.SkyblockItemRecipe, error) {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("recipes")

	var recipe structs.SkyblockItemRecipe
	
	filter := bson.M{"id": id}
	opts := options.FindOne().SetProjection(bson.M{"_id": 0})
	err := collection.FindOne(
		context.TODO(),
		filter,
		opts,
	).Decode(&recipe)

	return recipe, err
}