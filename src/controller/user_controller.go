package controller

import (
	"github.com/gin-gonic/gin"

	userapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/user"
)

func NewUserControllerInterface(
	serviceInterface userapp.UserDomainService,
) UserControllerInterface {
	return &userControllerInterface{
		service: serviceInterface,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindUserById(c *gin.Context)
	FindUserByEmail(c *gin.Context)
	UpdateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type userControllerInterface struct {
	service userapp.UserDomainService
}
