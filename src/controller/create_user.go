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

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, errors=%s", err.Error())
		// logger.Error("Error trying to validate user info", err, zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Name,
		userRequest.Password,
		userRequest.Age,
	)

	domainResult, err := uc.service.CreateUserService(domain)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainResult))
}
