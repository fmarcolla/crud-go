package test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindUserByEmail(t *testing.T) {
	t.Run("user_not_found_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test11@test.com",
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_success", func(t *testing.T) {
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

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: email,
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(ctx)

		var response map[string]string
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		Id_response, _ := response["id"]
		Email_response, _ := response["email"]

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, id, Id_response)
		assert.EqualValues(t, email, Email_response)
	})
}

func TestFindUserById(t *testing.T) {
	t.Run("user_not_found_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		id := primitive.NewObjectID()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		UserController.FindUserById(context)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		id := primitive.NewObjectID()
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

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserById(ctx)

		var response map[string]string
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		Id_response, _ := response["id"]
		Email_response, _ := response["email"]

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, id.Hex(), Id_response)
		assert.EqualValues(t, email, Email_response)
	})
}
