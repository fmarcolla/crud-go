package service

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
)

func (ud *userDomainService) UserLoginService(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *rest_err.RestErr) {
	userLogin, _ := ud.userRepository.FindUserByEmail(userDomain.GetEmail())

	if userLogin == nil {
		return nil, "", rest_err.NewForbiddenError("Credentials does not match.")
	}

	passwordHash := userLogin.GetPassword()
	if userDomain.CheckPasswordHash(passwordHash) == false {
		return nil, "", rest_err.NewForbiddenError("Credentials does not match.")
	}

	token, err := userLogin.GenerateToken()

	if err != nil {
		return nil, "", err
	}

	return userLogin, token, nil
}
