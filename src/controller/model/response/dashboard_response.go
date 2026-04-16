package response

import dashboarddomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/dashboard"

type DashboardMetricsResponse struct {
	Metrics *dashboarddomain.DashboardMetrics `json:"metrics"`
}
