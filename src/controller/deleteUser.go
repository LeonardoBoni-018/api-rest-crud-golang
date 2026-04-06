package controller

import (
	"net/http"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	loger.Info("Init DeleteUser controller",
		zap.String("journey", "DeleteUser"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		errRest := rest_err.NewBadRequestError("UserId is not valid hex")
		c.JSON(errRest.Code, errRest)
		return
	}

	if err := uc.service.DeleteUserDomain(userId); err != nil {
		logger.Info("Error trying to call DeleteUser", zap.String("journey", "DeleteUser"))

		c.JSON(err.Code, err)
		return
	}

	loger.Info("User Updated successfully",
		zap.String("journey", "DeleteUser"),
	)
	c.Status(http.StatusOK)
}

