package service

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
	"crud-go/src/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_CreateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)
	t.Run("when_user_already_exists_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "E-mail already exist.")
	})

	t.Run("when_user_is_not_registered_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repository.EXPECT().CreateUser(userDomain).Return(nil, rest_err.NewInternalServerError("error on trying to create a user"))

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error on trying to create a user")
	})

	t.Run("when_user_is_not_registered_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(nil, nil)
		repository.EXPECT().CreateUser(userDomain).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.NotNil(t, user)
		assert.Nil(t, err)
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
	})
}
