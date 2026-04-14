package booking

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	repository "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository"

)

type BookingService interface {
	CreateBooking(*domain.Booking) (*domain.Booking, *rest_err.RestErr)
	FindBookingsByTenantID(string) ([]*domain.Booking, *rest_err.RestErr)
	FindBookingByID(string) (*domain.Booking, *rest_err.RestErr)
}

type bookingService struct {
	bookingRepository repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) BookingService {
	return &bookingService{
		bookingRepository: repo,
	}
}
