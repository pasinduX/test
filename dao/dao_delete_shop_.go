package dao

import (
	"TestApp/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteShop (objectid string)  error {

    result, err := dbConfig.DATABASE.Collection("Shops").DeleteOne(context.Background(), bson.M{"objectid": objectid})
    if err != nil {
    	return err
    }
    if result.DeletedCount < 1 {
    	return errors.New("Specified Id not found!")
    }
    return nil
}