package repository

import (
	"context"
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
	"crud-go/src/model/repository/entity/converter"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr {
	collection := ur.databaseConnection.Collection("users")

	userEntity := converter.ConvertDomainToEntity(userDomain)

	objectId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{{Key: "$set", Value: userEntity}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		rest_err.NewInternalServerError(err.Error())
	}

	return nil
}
