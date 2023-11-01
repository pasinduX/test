package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallGroceryItems () (*[]dto.GroceryItems, error) {
	var objects []dto.GroceryItems
	results, err := dbConfig.DATABASE.Collection("GroceryItemss").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching GroceryItems")
	}
	for results.Next(context.Background()) {
		var object dto.GroceryItems
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding GroceryItems")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
