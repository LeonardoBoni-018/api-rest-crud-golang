package booking

import (
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	bookingdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	servicedomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
)

func (s *bookingService) GetServicesByTenantSlug(slug string) ([]*servicedomain.Service, *rest_err.RestErr) {
	tenant, err := s.tenantRepository.FindTenantBySlug(slug)
	if err != nil {
		return nil, err
	}

	return s.serviceRepository.FindServicesByTenantID(tenant.ID)
}

func (s *bookingService) GetAvailabilityByTenantSlug(slug, serviceID string, date time.Time) ([]string, *rest_err.RestErr) {
	tenant, err := s.tenantRepository.FindTenantBySlug(slug)
	if err != nil {
		return nil, err
	}

	return s.GetAvailableSlots(tenant.ID, serviceID, date)
}

func (s *bookingService) CreateBookingForTenantSlug(slug string, booking *bookingdomain.Booking) (*bookingdomain.Booking, *rest_err.RestErr) {
	tenant, err := s.tenantRepository.FindTenantBySlug(slug)
	if err != nil {
		return nil, err
	}

	service, err := s.serviceRepository.FindServiceByID(booking.ServiceID)
	if err != nil {
		return nil, err
	}
	if service.TenantID != tenant.ID {
		return nil, rest_err.NewBadRequestError("service does not belong to tenant")
	}

	booking.TenantID = tenant.ID
	if booking.Status == "" {
		booking.Status = "pending"
	}

	return s.bookingRepository.CreateBooking(booking)
}

func (s *bookingService) UpdateBookingStatus(tenantID, bookingID, status string) (*bookingdomain.Booking, *rest_err.RestErr) {
	booking, err := s.bookingRepository.FindBookingByID(bookingID)
	if err != nil {
		return nil, err
	}

	if booking.TenantID != tenantID {
		return nil, rest_err.NewForbiddenError("you are not authorized to update this booking")
	}

	booking.Status = status
	booking.UpdatedAt = time.Now()
	return s.bookingRepository.UpdateBooking(booking)
}
