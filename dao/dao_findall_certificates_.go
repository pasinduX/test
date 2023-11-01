package dao

import (
	"TestApp/dbConfig"
	"TestApp/dto"
	"go.mongodb.org/mongo-driver/bson"
    "context"
    "errors"
)

func DB_FindallCertificates () (*[]dto.Certificates, error) {
	var objects []dto.Certificates
	results, err := dbConfig.DATABASE.Collection("Certificatess").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.New("Error When Fetching Certificates")
	}
	for results.Next(context.Background()) {
		var object dto.Certificates
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding Certificates")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
