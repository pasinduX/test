package dao

import (
	"TestApp/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteGroceryItems (objectid string)  error {

    result, err := dbConfig.DATABASE.Collection("GroceryItemss").DeleteOne(context.Background(), bson.M{"objectid": objectid})
    if err != nil {
    	return err
    }
    if result.DeletedCount < 1 {
    	return errors.New("Specified Id not found!")
    }
    return nil
}