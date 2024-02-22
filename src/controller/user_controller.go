package controller

import (
	"crud-go/src/model/service"

	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(serviceInterface service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	LoginUser(c *gin.Context)

	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
