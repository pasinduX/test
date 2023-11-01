package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"context"
	"errors"
    "go.mongodb.org/mongo-driver/bson"
)

func DB_FindSupplierbyObjectId (objectid string) (*dto.Supplier, error) {
	var object dto.Supplier

	err := dbConfig.DATABASE.Collection("Suppliers").FindOne(context.Background(), bson.M{"objectid": objectid}).Decode(&object)
	if err != nil {
    	return nil, errors.New("Error When Fetching Supplier")
    }
	if object == (dto.Supplier{}) {
		return nil, errors.New("Specified ID not found!")
	}
	return &object, nil
}
