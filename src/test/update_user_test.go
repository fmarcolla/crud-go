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

func TestUpdateUser(t *testing.T) {
	t.Run("Should be able to update an user", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		id := primitive.NewObjectID()
		email := generateRandomEmail()

		newUser := bson.M{
			"_id":   id,
			"name":  "Old name",
			"email": email,
			"age":   15,
		}
		_, err := Database.Collection("users").InsertOne(context.Background(), newUser)

		if err != nil {
			t.Fatal(err)
			return
		}

		userRequest := request.UserRequest{
			Name: "New Name",
			Age:  25,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		MakeRequest(ctx, param, url.Values{}, "PUT", stringReader)
		UserController.UpdateUser(ctx)

		filter := bson.D{{Key: "_id", Value: id}}
		userEntity := &entity.UserEntity{}

		Database.Collection("users").FindOne(context.Background(), filter).Decode(userEntity)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, userEntity.Name, userRequest.Name)
		assert.EqualValues(t, userEntity.Age, userRequest.Age)
	})
}
