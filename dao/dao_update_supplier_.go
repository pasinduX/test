package dao

import (
	"TestApp/dbConfig"
    "TestApp/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateSupplier (object *dto.Supplier)  error {

	result, err := dbConfig.DATABASE.Collection("Suppliers").UpdateOne(context.Background(), bson.M{"objectid": object.ObjectId}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}