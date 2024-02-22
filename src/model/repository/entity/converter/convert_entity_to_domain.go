package converter

import (
	"crud-go/src/model"
	"crud-go/src/model/repository/entity"
)

func ConvertEntityToDomain(userEntity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		userEntity.Email,
		userEntity.Name,
		userEntity.Password,
		userEntity.Age,
	)

	domain.SetID(userEntity.ID.Hex())

	return domain
}
