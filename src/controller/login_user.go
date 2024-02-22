package controller

import (
	"crud-go/src/configuration/validation"
	"crud-go/src/controller/model/request"
	"crud-go/src/model"
	"crud-go/src/view"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	var userLoginRequest request.UserLoginRequest

	if err := c.ShouldBindJSON(&userLoginRequest); err != nil {
		log.Printf("Error trying to marshal object, errors=%s", err.Error())
		// logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(
		userLoginRequest.Email,
		userLoginRequest.Password,
	)

	domainResult, token, err := uc.service.UserLoginService(domain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
