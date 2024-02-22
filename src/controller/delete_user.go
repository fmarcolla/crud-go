package controller

import (
	"crud-go/src/configuration/rest_err"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	_, errId := primitive.ObjectIDFromHex(userId)

	if errId != nil {
		errRest := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(errRest.Code, errRest)
		return
	}

	err := uc.service.DeleteUserService(userId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
