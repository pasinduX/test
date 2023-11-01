package dao

import (
    "context"
	"TestApp/dbConfig"
	"TestApp/dto"

)

func DB_CreateCertificates (application *dto.Certificates) error {

	_, err := dbConfig.DATABASE.Collection("Certificatess").InsertOne(context.Background(), application)
	if err != nil {
		return err
	}
	return nil
}