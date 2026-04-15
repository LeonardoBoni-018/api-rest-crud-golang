package booking

import (
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	bookingRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository"
	repository "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository"
	serviceRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository"
	tenantRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
)

type BookingService interface {
	CreateBooking(*domain.Booking) (*domain.Booking, *rest_err.RestErr)
	FindBookingsByTenantID(string) ([]*domain.Booking, *rest_err.RestErr)
	FindBookingByID(string) (*domain.Booking, *rest_err.RestErr)
	GetAvailableSlots(tenantID, serviceID string, date time.Time) ([]string, *rest_err.RestErr)
}

type bookingService struct {
	bookingRepository repository.BookingRepository
	tenantRepository  tenantRepo.TenantRepository
	serviceRepository serviceRepo.ServiceRepository
}

func NewBookingService(
	bookingRepo bookingRepo.BookingRepository,
	tenantRepo tenantRepo.TenantRepository,
	serviceRepo serviceRepo.ServiceRepository,
) BookingService {
	return &bookingService{
		bookingRepository: bookingRepo,
		tenantRepository:  tenantRepo,
		serviceRepository: serviceRepo,
	}
}
