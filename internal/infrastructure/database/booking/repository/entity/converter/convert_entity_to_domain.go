package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository/entity"
)

func ToDomain(entity entity.BookingEntity) *booking.Booking {
	return &booking.Booking{
		ID:            entity.ID.Hex(),
		TenantID:      entity.TenantID,
		ServiceID:     entity.ServiceID,
		CustomerID:    entity.CustomerID,
		Date:          entity.Date,
		TimeSlot:      entity.TimeSlot,
		Duration:      entity.Duration,
		Status:        entity.Status,
		CustomerName:  entity.CustomerName,
		CustomerPhone: entity.CustomerPhone,
		CustomerEmail: entity.CustomerEmail,
		Notes:         entity.Notes,
		CreatedAt:     entity.CreatedAt.Time(),
		UpdatedAt:     entity.UpdatedAt.Time(),
	}
}
