package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallSupplier () (*[]dto.Supplier, error) {
	var objects []dto.Supplier
	results, err := dbConfig.DATABASE.Collection("Suppliers").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Supplier")
	}
	for results.Next(context.Background()) {
		var object dto.Supplier
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Supplier")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
