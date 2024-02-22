package service

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {
	user, _ := ud.userRepository.FindUserByEmail(userDomain.GetEmail())

	if user != nil {
		return nil, rest_err.NewBadRequestError("E-mail already exist.")
	}

	userDomain.EncryptPassword()
	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)

	if err != nil {
		return nil, err
	}

	return userDomainRepository, nil
}
