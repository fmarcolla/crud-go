package controller

import (
	"crud-go/src/configuration/rest_err"
	"crud-go/src/view"
	"net/http"
	"net/mail"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errorMessage := "User Id invalid"
		err := rest_err.NewBadRequestError(errorMessage)

		c.JSON(err.Code, err)
		return
	}

	userDomain, err := uc.service.FindUserByIdService(userId)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	email := c.Param("userEmail")

	if _, err := mail.ParseAddress(email); err != nil {
		errorMessage := "User e-mail invalid"
		err := rest_err.NewBadRequestError(errorMessage)

		c.JSON(err.Code, err)
		return
	}

	userDomain, err := uc.service.FindUserByEmailService(email)

	if err != nil {
		c.JSON(err.Code, err)

		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
