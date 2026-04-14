package response

import domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"

type BookingResponse struct {
	Booking *domain.Booking `json:"booking"`
}

type BookingsResponse struct {
	Bookings []*domain.Booking `json:"bookings"`
}
