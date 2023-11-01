package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallShop () (*[]dto.Shop, error) {
	var objects []dto.Shop
	results, err := dbConfig.DATABASE.Collection("Shops").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Shop")
	}
	for results.Next(context.Background()) {
		var object dto.Shop
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Shop")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
