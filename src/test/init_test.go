package test

import (
	"bytes"
	"crud-go/src/controller"
	"crud-go/src/model/repository"
	"crud-go/src/model/service"
	"crud-go/src/test/connection"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

func TestMain(m *testing.M) {
	err := os.Setenv("MONGODB_NAME", "crud-go-test")
	if err != nil {
		return
	}

	closeConnection := func() {}
	Database, closeConnection = connection.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service.NewUserDomainService(repo)
	UserController = controller.NewUserControllerInterface(userService)

	defer func() {
		os.Clearenv()
		closeConnection()
	}()

	os.Exit(m.Run())
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func generateRandomEmail() string {
	randomString := generateRandomString(10)

	var buffer bytes.Buffer

	buffer.WriteString(randomString)
	buffer.WriteString("@test.com")

	return buffer.String()
}
