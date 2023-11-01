package dao

import (
    "context"
	"TestApp/dbConfig"
	"TestApp/dto"

)

func DB_CreateSupplier (application *dto.Supplier) error {

	_, err := dbConfig.DATABASE.Collection("Suppliers").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}