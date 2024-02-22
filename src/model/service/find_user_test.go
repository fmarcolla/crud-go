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

func TestUserDomainService_FindUserByIdService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)

		repository.EXPECT().FindUserById(id).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByIdService(id)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
	})

	t.Run("when_does_not_exists_an_user_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().FindUserById(id).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByIdService(id)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainReturn)
		assert.EqualValues(t, err.Message, "user not found")
	})
}

func TestUserDomainService_FindUserByEmailService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_exists_an_user_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()
		email := "test@test.com"

		userDomain := model.NewUserDomain(email, "test", "test", 10)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(email).Return(userDomain, nil)

		userDomainReturn, err := service.FindUserByEmailService(email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomainReturn.GetID(), id)
		assert.EqualValues(t, userDomainReturn.GetEmail(), userDomain.GetEmail())
	})

	t.Run("when_does_not_exists_an_user_return_error", func(t *testing.T) {
		email := "test@test.com"

		repository.EXPECT().FindUserByEmail(email).Return(nil, rest_err.NewNotFoundError("user not found"))
		userDomainReturn, err := service.FindUserByEmailService(email)

		assert.NotNil(t, err)
		assert.Nil(t, userDomainReturn)
		assert.EqualValues(t, err.Message, "user not found")
	})
}
