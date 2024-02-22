package service

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
	"crud-go/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserService(model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUserService(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIdService(string) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailService(string) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserService(string) *rest_err.RestErr
	UserLoginService(model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr)
}
