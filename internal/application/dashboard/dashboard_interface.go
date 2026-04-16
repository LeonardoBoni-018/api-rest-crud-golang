package dashboard

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	dashboarddomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/dashboard"
	bookingRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository"
	serviceRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository"
)

type DashboardService interface {
	GetMetrics(tenantID string) (*dashboarddomain.DashboardMetrics, *rest_err.RestErr)
}

type dashboardService struct {
	bookingRepository bookingRepo.BookingRepository
	serviceRepository serviceRepo.ServiceRepository
}

func NewDashboardService(
	bookingRepository bookingRepo.BookingRepository,
	serviceRepository serviceRepo.ServiceRepository,
) DashboardService {
	return &dashboardService{
		bookingRepository: bookingRepository,
		serviceRepository: serviceRepository,
	}
}
