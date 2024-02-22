package test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDeleteUser(t *testing.T) {
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

	MakeRequest(ctx, param, url.Values{}, "DELETE", nil)
	UserController.DeleteUser(ctx)

	filter := bson.D{{Key: "_id", Value: id}}
	result := Database.Collection("users").FindOne(context.Background(), filter)

	assert.EqualValues(t, http.StatusOK, recorder.Code)
	assert.NotNil(t, result.Err())
}
