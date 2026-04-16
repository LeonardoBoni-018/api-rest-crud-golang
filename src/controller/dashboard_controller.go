package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dashboardapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/dashboard"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/response"
)

type DashboardControllerInterface interface {
	GetMetrics(c *gin.Context)
}

type dashboardController struct {
	dashboard dashboardapp.DashboardService
}

func NewDashboardController(service dashboardapp.DashboardService) DashboardControllerInterface {
	return &dashboardController{dashboard: service}
}

func (dc *dashboardController) GetMetrics(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	metrics, err := dc.dashboard.GetMetrics(tenantID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response.DashboardMetricsResponse{Metrics: metrics})
}
