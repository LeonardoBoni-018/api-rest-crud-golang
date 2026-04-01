package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/view"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	loger.Info("Init findUserById",
		zap.String("journey", "findUserById"))

	// user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	// if err != nil {
	// 	c.JSON(err.Code, err)
	// 	return
	// }
	// logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		loger.Error("Error trying to validate id", err,
			zap.String("journey", "findUserById"))
		errorMessage := rest_err.NewBadRequestError(
			"UserId is not valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIdServices(userId)
	if err != nil {
		loger.Error("Error trying to call findUserByid servvices ", err,
			zap.String("journey", "findUserById"))
		c.JSON(err.Code, err)
		return
	}

	loger.Info("findUserById controller executed succesfully",
		zap.String("journey", "findUserById"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	loger.Info("Init findUserByEmail",
		zap.String("journey", "findUserByEmail"))

	// user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	// if err != nil {
	// 	c.JSON(err.Code, err)
	// 	return
	// }
	// logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {
		loger.Error("Error trying to validate email", err,
			zap.String("journey", "findUserByEmail"))
		errorMessage := rest_err.NewBadRequestError(
			"UserEmail is not valid id")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		loger.Error("Error trying to call findUserByEmail servvices ", err,
			zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	loger.Info("findUserByEmail controller executed succesfully",
		zap.String("journey", "findUserByEmail"))
	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}
