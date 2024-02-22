package test

import (
	"context"
	"crud-go/src/controller/model/request"
	"crud-go/src/model/repository/entity"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateUser(t *testing.T) {
	t.Run("user_already_exists_with_email_return_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		id := primitive.NewObjectID().Hex()
		email := generateRandomEmail()

		newUser := bson.M{
			"_id":   id,
			"name":  "test",
			"email": email,
		}
		_, err := Database.Collection("users").InsertOne(context.Background(), newUser)

		if err != nil {
			t.Fatal(err)
			return
		}

		userRequest := request.UserRequest{
			Email:    email,
			Password: "test!@#123",
			Name:     "Test",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("create_user_return_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		email := generateRandomEmail()

		userRequest := request.UserRequest{
			Email:    email,
			Password: "test!@#123",
			Name:     "Test",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(ctx, []gin.Param{}, url.Values{}, "POST", stringReader)
		UserController.CreateUser(ctx)

		filter := bson.D{{Key: "email", Value: email}}
		userEntity := &entity.UserEntity{}

		Database.Collection("users").FindOne(context.Background(), filter).Decode(userEntity)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, userEntity.Email, userRequest.Email)
		assert.EqualValues(t, userEntity.Name, userRequest.Name)
		assert.EqualValues(t, userEntity.Age, userRequest.Age)
	})
}
