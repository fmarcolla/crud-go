package controller

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/configuration/validation"
	"crud-go/src/controller/model/request"
	"crud-go/src/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	var updateUserRequest request.UpdateUserRequest

	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		log.Printf("Error trying to marshal object, errors=%s", err.Error())
		// logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	userId := c.Param("userId")
	_, errId := primitive.ObjectIDFromHex(userId)

	if errId != nil {
		errRest := rest_err.NewBadRequestError("Invalid user id")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(
		updateUserRequest.Name,
		updateUserRequest.Age,
	)

	err := uc.service.UpdateUserService(userId, domain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Status(http.StatusOK)
}
