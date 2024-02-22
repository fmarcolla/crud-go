package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_DeleteUser(t *testing.T) {
	database_name := "user_database_test"

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mTestDb.Close()

	mTestDb.Run("when_delete_user_return_success", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "n", Value: 1},
			{Key: "acknowledged", Value: true},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		err := repo.DeleteUser("test")

		assert.Nil(t, err)
	})

	// mTestDb.Run("return_error_from_database", func(mt *mtest.T) {
	// 	mt.AddMockResponses(bson.D{
	// 		{Key: "ok", Value: 1},
	// 		{Key: "n", Value: 0},
	// 		{Key: "acknowledged", Value: true},
	// 	})

	// 	databaseMock := mt.Client.Database(database_name)

	// 	repo := NewUserRepository(databaseMock)
	// 	err := repo.DeleteUser("test")

	// 	assert.NotNil(t, err)
	// })
}
