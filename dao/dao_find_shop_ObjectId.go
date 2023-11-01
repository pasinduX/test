package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindShopbyObjectId (objectid string) (*dto.Shop, error) {
	var object dto.Shop

	err := dbConfig.DATABASE.Collection("Shops").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Shop")
    }
	if object == (dto.Shop{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
