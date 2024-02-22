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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserControllerInterface_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)
	t.Run("validation_got_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLoginRequest{
			Email:    "TEST_ERROR",
			Password: "test",
		}

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("validation_is_valid_but_services_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLoginRequest{
			Email:    "test@test.com",
			Password: "test!@#123",
		}

		domain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		service.EXPECT().UserLoginService(domain).Return(nil, "", rest_err.NewInternalServerError("error test"))

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})
	t.Run("validation_is_valid_and_services_returns_ok", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		userRequest := request.UserLoginRequest{
			Email:    "test@test.com",
			Password: "test!@#123",
		}

		domain := model.NewUserLoginDomain(
			userRequest.Email,
			userRequest.Password,
		)

		b, _ := json.Marshal(userRequest)
		stringReader := io.NopCloser(strings.NewReader(string(b)))

		token := primitive.NewObjectID().Hex()

		service.EXPECT().UserLoginService(domain).Return(domain, token, nil)

		MakeRequest(context, []gin.Param{}, url.Values{}, "POST", stringReader)
		controller.LoginUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
		assert.EqualValues(t, recorder.Header().Values("Authorization")[0], token)
	})
}
