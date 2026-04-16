package booking

import (
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	servicedomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
	bookingRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository"
	serviceRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository"
	tenantRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
)

type BookingService interface {
	CreateBooking(*domain.Booking) (*domain.Booking, *rest_err.RestErr)
	FindBookingsByTenantID(string) ([]*domain.Booking, *rest_err.RestErr)
	FindBookingByID(string) (*domain.Booking, *rest_err.RestErr)
	GetAvailableSlots(tenantID, serviceID string, date time.Time) ([]string, *rest_err.RestErr)
	GetServicesByTenantSlug(string) ([]*servicedomain.Service, *rest_err.RestErr)
	GetAvailabilityByTenantSlug(string, string, time.Time) ([]string, *rest_err.RestErr)
	CreateBookingForTenantSlug(string, *domain.Booking) (*domain.Booking, *rest_err.RestErr)
	UpdateBookingStatus(string, string, string) (*domain.Booking, *rest_err.RestErr)
}

type bookingService struct {
	bookingRepository bookingRepo.BookingRepository
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
