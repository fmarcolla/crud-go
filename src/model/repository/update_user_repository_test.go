package repository

import (
	"crud-go/src/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_UpdateUser(t *testing.T) {
	database_name := "user_database_test"

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mTestDb.Close()

	mTestDb.Run("when_update_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain("test@teste.com", "test", "123456!@", 25)
		domain.SetID(primitive.NewObjectID().Hex())

		err := repo.UpdateUser(domain.GetID(), domain)

		assert.Nil(t, err)
	})

	// mTestDb.Run("return_error_from_database", func(mt *mtest.T) {
	// 	mt.AddMockResponses(bson.D{
	// 		{Key: "ok", Value: 0},
	// 		{Key: "n", Value: 0},
	// 		{Key: "acknowledged", Value: true},
	// 	})

	// 	databaseMock := mt.Client.Database(database_name)

	// 	repo := NewUserRepository(databaseMock)
	// 	domain := model.NewUserDomain("test@teste.com", "test", "123456!@", 25)
	// 	domain.SetID(primitive.NewObjectID().Hex())

	// 	err := repo.UpdateUser("asd", domain)

	// 	assert.NotNil(t, err)
	// })
}
