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
	cursor, err := collection.Find(
		context.TODO(),
		filter,
	)
	if err != nil {
		return nil, err
	}
	if err := cursor.All(context.Background(), &recipes); err != nil {
		return nil, err
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
	
	filter := bson.M{"itemId": id}
	err := collection.FindOne(
		context.TODO(),
		filter,
	).Decode(&recipe)

	return recipe, err
}

func UpdateRecipe(id string, data structs.SkyblockItemRecipe) error {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("recipes")

	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"itemId": id},
		bson.M{
			"$set": data,
		},
	)
	return err
}

func DeleteRecipe(id string) error {
	database := MongoClient.Database(constants.MongoDatabase)
	collection := database.Collection("recipes")

	_, err := collection.DeleteOne(
		context.TODO(),
		bson.M{"itemId": id},
	)
	return err
}