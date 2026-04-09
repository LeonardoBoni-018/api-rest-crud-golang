package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	serviceapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/service"
	serviceDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/response"
)

type ServiceControllerInterface interface {
	CreateService(c *gin.Context)
	ListServices(c *gin.Context)
	UpdateService(c *gin.Context)
	DeleteService(c *gin.Context)
}

type serviceController struct {
	service serviceapp.ServiceService
}

func NewServiceController(service serviceapp.ServiceService) ServiceControllerInterface {
	return &serviceController{service: service}
}

func (sc *serviceController) ListServices(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	services, err := sc.service.FindServicesByTenantID(tenantID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response.ServicesResponse{Services: services})
}

func (sc *serviceController) CreateService(c *gin.Context) {
	var req request.ServiceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
	}

	tenantID := c.GetString("tenant_id")
	domain := &serviceDomain.Service{
		TenantID:    tenantID,
		Name:        req.Name,
		Description: req.Description,
		Duration:    req.Duration,
		Price:       req.Price,
		IsActive:    true,
	}

	created, err := sc.service.CreateService(domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusCreated, response.ServiceResponse{Service: created})
}

func (sc *serviceController) UpdateService(c *gin.Context) {
	serviceID := c.Param("serviceId")
	var req request.ServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	tenantID := c.GetString("tenant_id")
	domain := &serviceDomain.Service{
		TenantID:    tenantID,
		Name:        req.Name,
		Description: req.Description,
		Duration:    req.Duration,
		Price:       req.Price,
		IsActive:    req.IsActive,
	}

	updated, err := sc.service.UpdateService(serviceID, domain)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response.ServiceResponse{Service: updated})
}

func (sc *serviceController) DeleteService(c *gin.Context) {
	serviceID := c.Param("serviceId")
	if err := sc.service.DeleteService(serviceID); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.Status(http.StatusNoContent)
}
