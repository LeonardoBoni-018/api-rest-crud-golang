package controller

import (
	"context"
	"net/http"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model/service"
	"github.com/gin-gonic/gin"
)

var businessService *service.BusinessService

func SetBusinessService(s *service.BusinessService) {
	businessService = s
}

func errorResponse(message string) gin.H {
	return gin.H{"error": message}
}

func CreateBusiness(c *gin.Context) {
	var req model.BusinessDomain
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}
	res, err := businessService.Create(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": res.InsertedID})
}
