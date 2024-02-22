package service

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/test/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserDomainService_DeleteUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_delete_an_user_return_success", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(nil)

		err := service.DeleteUserService(id)

		assert.Nil(t, err)
	})

	t.Run("when_delete_an_user_return_error", func(t *testing.T) {
		id := primitive.NewObjectID().Hex()

		repository.EXPECT().DeleteUser(id).Return(rest_err.NewInternalServerError("Error on trying to delete user"))

		err := service.DeleteUserService(id)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Error on trying to delete user")
	})
}
