package controller

import (
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/configuration/validation"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/view"
)

var (
	UserDomainInterface user.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	domain := user.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	if err := uc.service.CreateUserServices(domain); err != nil {
		logger.Info("Error trying to call CreateUser", zap.String("journey", "createUser"))

		c.JSON(err.Code, err)
		return
	}

	loger.Info("User created successfully",
		zap.String("journey", "CreateUser"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))
}


