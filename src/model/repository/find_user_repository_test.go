package repository

import (
	"crud-go/src/model/repository/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestUserRepository_FindUserByEmail(t *testing.T) {
	database_name := "user_database_test"

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mTestDb.Close()

	mTestDb.Run("when_sending_a_valid_email_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "123456!@",
			Name:     "Test",
			Age:      19,
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				fmt.Sprintf("%s.%s", database_name, "users"),
				mtest.FirstBatch,
				ConvertEntityToBSON(userEntity)),
		)

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail(userEntity.Email)

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
	})

	mTestDb.Run("return_error_mongodb", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.Nil(t, userDomain)
		assert.NotNil(t, err)
	})

	mTestDb.Run("return_error_document_not_found", func(mt *mtest.T) {
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				0,
				fmt.Sprintf("%s.%s", database_name, "users"),
				mtest.FirstBatch,
			))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserByEmail("test")

		assert.Nil(t, userDomain)
		assert.NotNil(t, err)
	})
}

func TestUserRepository_FindUserById(t *testing.T) {
	database_name := "user_database_test"

	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mTestDb.Close()

	mTestDb.Run("when_sending_a_valid_id_return_success", func(mt *mtest.T) {
		userEntity := entity.UserEntity{
			ID:       primitive.NewObjectID(),
			Email:    "test@test.com",
			Password: "123456!@",
			Name:     "Test",
			Age:      19,
		}

		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1,
				fmt.Sprintf("%s.%s", database_name, "users"),
				mtest.FirstBatch,
				ConvertEntityToBSON(userEntity)),
		)

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById(userEntity.ID.Hex())

		assert.Nil(t, err)
		assert.EqualValues(t, userDomain.GetID(), userEntity.ID.Hex())
		assert.EqualValues(t, userDomain.GetEmail(), userEntity.Email)
	})

	mTestDb.Run("return_error_mongodb", func(mt *mtest.T) {
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 0},
		})

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("test")

		assert.Nil(t, userDomain)
		assert.NotNil(t, err)
	})

	mTestDb.Run("return_error_document_not_found", func(mt *mtest.T) {
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				0,
				fmt.Sprintf("%s.%s", database_name, "users"),
				mtest.FirstBatch,
			))

		databaseMock := mt.Client.Database(database_name)

		repo := NewUserRepository(databaseMock)
		userDomain, err := repo.FindUserById("test")

		assert.Nil(t, userDomain)
		assert.Equal(t, err.Message, fmt.Sprintf("User not found with this ID test"))
		assert.NotNil(t, err)
	})
}

func ConvertEntityToBSON(e entity.UserEntity) bson.D {
	return bson.D{
		{Key: "_id", Value: e.ID},
		{Key: "email", Value: e.Email},
		{Key: "password", Value: e.Password},
		{Key: "name", Value: e.Name},
		{Key: "age", Value: e.Age},
	}
}
