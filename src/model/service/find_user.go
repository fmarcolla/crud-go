package service

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
)

func (ud *userDomainService) FindUserByIdService(userId string) (model.UserDomainInterface, *rest_err.RestErr) {
	userDomain, err := ud.userRepository.FindUserById(userId)

	if err != nil {
		return nil, err
	}

	return userDomain, nil
}

func (ud *userDomainService) FindUserByEmailService(userEmail string) (model.UserDomainInterface, *rest_err.RestErr) {
	userDomain, err := ud.userRepository.FindUserByEmail(userEmail)

	if err != nil {
		return nil, err
	}

	return userDomain, nil
}
