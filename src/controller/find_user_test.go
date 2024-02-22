package controller

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/model"
	"crud-go/src/test/mocks"
	"encoding/json"

	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)
	t.Run("email_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "TEST_ERROR",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		email := "test@test.com"

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: email,
			},
		}

		userDomain := model.NewUserDomain(email, "test", "123", 10)
		service.EXPECT().FindUserByEmailService(email).Return(userDomain, nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		var response map[string]string
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		Email_response, _ := response["email"]

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, "test@test.com", Email_response)
	})
}
func TestUserControllerInterface_FindUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)
	t.Run("id_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "TEST_ERROR",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserById(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIdService(id).Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserById(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		userDomain := model.NewUserDomain("test@test.com", "test", "123", 10)
		userDomain.SetID(id)

		service.EXPECT().FindUserByIdService(id).Return(userDomain, nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserById(context)

		var response map[string]string
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		Id_response, _ := response["id"]
		Email_response, _ := response["email"]

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, id, Id_response)
		assert.EqualValues(t, "test@test.com", Email_response)
	})
}
