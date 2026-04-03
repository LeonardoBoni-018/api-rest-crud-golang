package controller

import (
	"context"
	"net/http"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model/service"
	"github.com/gin-gonic/gin"
)

var serviceService *service.ServiceService

func SetServiceService(s *service.ServiceService) {
	serviceService = s
}

func CreateService(c *gin.Context) {
	businessID := c.Param("businessId")
	var req model.ServiceDomain
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	req.BusinessID = businessID

	res, err := serviceService.Create(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": res.InsertedID})
}

func GetServicesByBusiness(c *gin.Context) {
	businessID := c.Param("businessId")
	services, err := serviceService.FindByBusinessID(context.Background(), businessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, services)
}
