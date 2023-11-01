package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindGroceryItemsbyObjectId (objectid string) (*dto.GroceryItems, error) {
	var object dto.GroceryItems

	err := dbConfig.DATABASE.Collection("GroceryItemss").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching GroceryItems")
    }
	if object == (dto.GroceryItems{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
