package service

import "crud-go/src/configuration/rest_err"

func (ud *userDomainService) DeleteUserService(userId string) *rest_err.RestErr {
	err := ud.userRepository.DeleteUser(userId)

	if err != nil {
		return err
	}

	return nil
}
