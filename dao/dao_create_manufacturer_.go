package dao

import (
    "context"
	"TestApp/dbConfig"
	"TestApp/dto"

)

func DB_CreateManufacturer (application *dto.Manufacturer) error {

	_, err := dbConfig.DATABASE.Collection("Manufacturers").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}