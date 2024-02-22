package repository

import (
	"context"
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
	"crud-go/src/model/repository/entity"
	"crud-go/src/model/repository/entity/converter"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.databaseConnection.Collection("users")

	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this e-mail %s", email)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error on trying to find user by e-mail"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserById(id string) (model.UserDomainInterface, *rest_err.RestErr) {
	collection := ur.databaseConnection.Collection("users")

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	userEntity := &entity.UserEntity{}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this ID %s", id)
			return nil, rest_err.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error on trying to find user by ID"
		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	return converter.ConvertEntityToDomain(*userEntity), nil
}
