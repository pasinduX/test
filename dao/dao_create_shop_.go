package dao

import (
    "context"
	"TestApp/dbConfig"
	"TestApp/dto"

)

func DB_CreateShop (application *dto.Shop) error {

	_, err := dbConfig.DATABASE.Collection("Shops").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}