package converter

import (
	"crud-go/src/model"
	"crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(userDomain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Name:     userDomain.GetName(),
		Age:      userDomain.GetAge(),
	}
}
