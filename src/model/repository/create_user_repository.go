package repository

import (
	"context"
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
	"crud-go/src/model/repository/entity/converter"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.databaseConnection.Collection("users")

	userEntity := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), userEntity)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	userEntity.ID = result.InsertedID.(primitive.ObjectID)

	return converter.ConvertEntityToDomain(*userEntity), nil
}
