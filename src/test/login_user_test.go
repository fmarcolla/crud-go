package test

import (
	"crud-go/src/controller/model/request"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	t.Run("Should be able to user make login", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		email := generateRandomEmail()
		password := "test!@#123"

		createUserRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test",
			Age:      10,
		}

		bCreate, _ := json.Marshal(createUserRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		loginUserRequest := request.UserLoginRequest{
			Email:    email,
			Password: password,
		}

		bLogin, _ := json.Marshal(loginUserRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusOK, recorderLoginUser.Code)
		assert.NotEmpty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})

	t.Run("Should not be able to user make login with incorret credentials", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		email := generateRandomEmail()
		password := "test!@#123"

		createUserRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test",
			Age:      10,
		}

		bCreate, _ := json.Marshal(createUserRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		loginUserRequest := request.UserLoginRequest{
			Email:    email,
			Password: "asdasdtest!@#123asdaada",
		}

		bLogin, _ := json.Marshal(loginUserRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusForbidden, recorderLoginUser.Code)
		assert.Empty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})
}
