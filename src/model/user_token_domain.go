package model

import (
	"crud-go/src/configuration/rest_err"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv("JWT_SECRET_KEY")

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("error trying to generate a jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv("JWT_SECRET_KEY")

	tokenValue := c.Request.Header.Get("Authorization")
	tokenValue = RemoveBearerPrefix(tokenValue)

	token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token")

		c.JSON(errRest.Code, err)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedRequestError("Invalid token")

		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	fmt.Println(fmt.Sprintf("usuario logado %s", claims["id"].(string)))
}

func RemoveBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}
