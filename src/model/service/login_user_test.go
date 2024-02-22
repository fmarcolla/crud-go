package service

import (
	"crud-go/src/model"
	"crud-go/src/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_UserLoginService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_login_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)
		userDomain.EncryptPassword()

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		userDomainLogin := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomainLogin.SetID(id)

		userLogged, token, err := service.UserLoginService(userDomainLogin)

		assert.Nil(t, err)
		assert.NotNil(t, userLogged)
		assert.NotNil(t, token)
		assert.EqualValues(t, userDomain.GetID(), id)
	})

	t.Run("when_user_does_not_exists_return_error", func(t *testing.T) {
		repository.EXPECT().FindUserByEmail("test@test.com").Return(nil, nil)

		userDomainLogin := model.NewUserDomain("test@test.com", "test", "test", 10)

		userLogged, token, err := service.UserLoginService(userDomainLogin)

		assert.NotNil(t, err)
		assert.Nil(t, userLogged)
		assert.Empty(t, token)
	})

	t.Run("when_password_is_wrong_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)
		userDomain.EncryptPassword()

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		userDomainLogin := model.NewUserDomain("test@test.com", "test", "error_password", 10)
		userDomainLogin.SetID(id)

		userLogged, token, err := service.UserLoginService(userDomainLogin)

		assert.NotNil(t, err)
		assert.Nil(t, userLogged)
		assert.Empty(t, token)
	})
}
