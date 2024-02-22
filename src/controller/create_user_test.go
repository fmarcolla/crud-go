package controller

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/controller/model/request"
	"crud-go/src/model"
	"crud-go/src/test/mocks"
	"encoding/json"
	"io"
	"strings"

	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)
	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "TEST_ERROR",
			Password: "test",
			Name:     "Test",
			Age:      10,
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_is_valid_but_services_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "test!@#123",
			Name:     "Test",
			Age:      10,
		}

		domain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Name,
			userRequest.Password,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateUserService(domain).Return(nil, rest_err.NewInternalServerError("error test"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})
	t.Run("validation_is_valid_and_services_returns_ok", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserRequest{
			Email:    "test@test.com",
			Password: "test!@#123",
			Name:     "Test",
			Age:      10,
		}

		domain := model.NewUserDomain(
			userRequest.Email,
			userRequest.Name,
			userRequest.Password,
			userRequest.Age,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().CreateUserService(domain).Return(domain, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.CreateUser(context)

		var response map[string]string
		_ = json.Unmarshal([]byte(recorder.Body.String()), &response)

		Id_response, _ := response["id"]
		Email_response, _ := response["email"]

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, "test@test.com", Email_response)
		assert.NotNil(t, Id_response)
	})
}
