package routes

import (
	"crud-go/src/controller"
	"crud-go/src/model"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/users/id/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/users/email/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.PUT("/users/:userId", model.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/users/:userId", model.VerifyTokenMiddleware, userController.DeleteUser)

	r.POST("/users", userController.CreateUser)
	r.POST("/login", userController.LoginUser)
}
