package repository

import (
	"crud-go/src/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_CreateUser(t *testing.T) {
	database_name := "user_database_test"

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mTestDb.Close()

	mTestDb.Run("when_create_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain("test@teste.com", "test", "123456!@", 25)

		userDomain, err := repo.CreateUser(domain)

		_, errId := primitive.ObjectIDFromHex(userDomain.GetID())

		assert.Nil(t, err)
		assert.Nil(t, errId)
		assert.EqualValues(t, userDomain.GetEmail(), domain.GetEmail())
		assert.EqualValues(t, userDomain.GetName(), domain.GetName())
	})

	mTestDb.Run("return_error_from_database", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		domain := model.NewUserDomain("test@teste.com", "test", "123456!@", 25)

		userDomain, err := repo.CreateUser(domain)

		assert.NotNil(t, err)
		assert.Nil(t, userDomain)
	})
}
