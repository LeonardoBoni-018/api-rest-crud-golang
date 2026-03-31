package controller

import (
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/configuration/validation"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/view"
)

func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	loger.Info("Init LoginUser controller",
		zap.String("journey", "LoginUser"),
	)
	var userRequest request.UserLogin

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		loger.Error("Error try to validate user info", err,
			zap.String("journey", "LoginUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserLoginDomain(userRequest.Email, userRequest.Password)
	domainUserResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Info("Error trying to call LoginUser", zap.String("journey", "LoginUser"))

		c.JSON(err.Code, err)
		return
	}

	loger.Info("User created successfully",
		zap.String("journey", "LoginUser"),
	)
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domainUserResult))
}
