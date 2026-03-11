package controller

import (
	"net/http"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/configuration/validation"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	loger.Info("Init CreateUser controller",
		zap.String("journey", "CreateUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		loger.Error("Error try to validate user info", err,
			zap.String("journey", "CreateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	service := service.NewUserDomainService()
	if err := service.CreateUserDomain(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// response := response.UserResponse{
	// 	ID:    "test",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }
	loger.Info("User created successfully",
		zap.String("journey", "CreateUser"),
	)
	c.String(http.StatusOK, "")
}
