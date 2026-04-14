package booking

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"

)

func (s *bookingService) CreateBooking(domain *domain.Booking) (*domain.Booking, *rest_err.RestErr) {
	return s.bookingRepository.CreateBooking(domain)
}
