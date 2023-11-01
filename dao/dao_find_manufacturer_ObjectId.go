package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindManufacturerbyObjectId (objectid string) (*dto.Manufacturer, error) {
	var object dto.Manufacturer

	err := dbConfig.DATABASE.Collection("Manufacturers").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Manufacturer")
    }
	if object == (dto.Manufacturer{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
