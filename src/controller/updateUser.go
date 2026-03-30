package controller

import (
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/configuration/validation"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	loger.Info("Init UpdateUser controller",
		zap.String("journey", "UpdateUser"),
	)
	var userRequest request.UserUpdateRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		loger.Error("Error try to validate user info", err,
			zap.String("journey", "UpdateUser"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("UserId is not valid hex")
		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserUpdateDomain(userRequest.Age, userRequest.Name)
	if err := uc.service.UpdateUserDomain(userId, domain); err != nil {
		logger.Info("Error trying to call UpdateUser", zap.String("journey", "UpdateUser"))

		c.JSON(err.Code, err)
		return
	}

	loger.Info("User Updated successfully",
		zap.String("journey", "UpdateUser"),
	)
	c.Status(http.StatusOK)
}
