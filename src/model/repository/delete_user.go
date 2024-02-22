package repository

import (
	"context"
	"crud-go/src/configuration/rest_err"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ur *userRepository) DeleteUser(userId string) *rest_err.RestErr {
	collection := ur.databaseConnection.Collection("users")

	objectId, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: objectId}}

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		rest_err.NewInternalServerError(err.Error())
	}

	if result.DeletedCount < 1 {
		rest_err.NewNotFoundError("Register not found!")
	}

	return nil
}
