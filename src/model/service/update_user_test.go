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

func TestUserDomainService_UpdateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_update_an_user_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(userDomain.GetID(), userDomain).Return(nil)

		err := service.UpdateUserService(userDomain.GetID(), userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), id)
	})

	t.Run("when_update_an_user_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 10)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(userDomain.GetID(), userDomain).Return(rest_err.NewInternalServerError("Error on trying to update user"))

		err := service.UpdateUserService(userDomain.GetID(), userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error on trying to update user")
	})
}
