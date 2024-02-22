package repository

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *rest_err.RestErr
	FindUserByEmail(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
