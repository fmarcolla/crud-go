package main

import (
	"crud-go/src/controller"
	"crud-go/src/model/repository"
	"crud-go/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	userController := controller.NewUserControllerInterface(service)

	return userController
}
