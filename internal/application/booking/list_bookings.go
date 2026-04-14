package booking

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
)

func (s *bookingService) FindBookingsByTenantID(tenantID string) ([]*domain.Booking, *rest_err.RestErr) {
	return s.bookingRepository.FindBookingsByTenantID(tenantID)
}

func (s *bookingService) FindBookingByID(id string) (*domain.Booking, *rest_err.RestErr) {
	return s.bookingRepository.FindBookingByID(id)
}
