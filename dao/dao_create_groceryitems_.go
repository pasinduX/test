package dao

import (
    "context"
	"TestApp/dbConfig"
	"TestApp/dto"

)

func DB_CreateGroceryItems (application *dto.GroceryItems) error {

	_, err := dbConfig.DATABASE.Collection("GroceryItemss").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}