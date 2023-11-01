package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallManufacturer () (*[]dto.Manufacturer, error) {
	var objects []dto.Manufacturer
	results, err := dbConfig.DATABASE.Collection("Manufacturers").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Manufacturer")
	}
	for results.Next(context.Background()) {
		var object dto.Manufacturer
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Manufacturer")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
