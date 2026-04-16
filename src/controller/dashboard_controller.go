package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	dashboardapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/dashboard"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/response"
)

type DashboardControllerInterface interface {
	GetMetrics(c *gin.Context)
	GetCalendar(c *gin.Context)
	GetReports(c *gin.Context)
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

func (dc *dashboardController) GetCalendar(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	calendar, err := dc.dashboard.GetCalendar(tenantID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response.DashboardCalendarResponse{Calendar: calendar})
}

func (dc *dashboardController) GetReports(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	reports, err := dc.dashboard.GetReports(tenantID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response.DashboardReportsResponse{Reports: reports})
}
