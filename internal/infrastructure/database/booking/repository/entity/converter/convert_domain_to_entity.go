package converter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository/entity"

)

func ToEntity(domain *booking.Booking) *entity.BookingEntity {
	entity := &entity.BookingEntity{
		TenantID:      domain.TenantID,
		ServiceID:     domain.ServiceID,
		CustomerID:    domain.CustomerID,
		Date:          domain.Date,
		TimeSlot:      domain.TimeSlot,
		Duration:      domain.Duration,
		Status:        domain.Status,
		CustomerName:  domain.CustomerName,
		CustomerPhone: domain.CustomerPhone,
		CustomerEmail: domain.CustomerEmail,
		Notes:         domain.Notes,
		CreatedAt:     primitive.NewDateTimeFromTime(domain.CreatedAt),
		UpdatedAt:     primitive.NewDateTimeFromTime(domain.UpdatedAt),
	}

	if domain.ID != "" {
		objectID, err := primitive.ObjectIDFromHex(domain.ID)
		if err == nil {
			entity.ID = objectID
		}
	}
	
	return entity
}