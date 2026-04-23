package booking

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID             string    `json:"id" bson:"_id"`
	TenantID       string    `json:"tenant_id" bson:"tenant_id"`
	ServiceID      string    `json:"service_id" bson:"service_id"`
	CustomerID     string    `json:"customer_id" bson:"customer_id"`
	Date           string    `json:"date" bson:"date"`
	TimeSlot       string    `json:"time_slot" bson:"time_slot"`
	Duration       int       `json:"duration" bson:"duration"`
	Status         string    `json:"status" bson:"status"`
	CustomerName   string    `json:"customer_name" bson:"customer_name"`
	CustomerPhone  string    `json:"customer_phone" bson:"customer_phone"`
	CustomerEmail string    `json:"customer_email" bson:"customer_email"`
	Notes          string    `json:"notes" bson:"notes"`
	CancelToken    string    `json:"cancel_token" bson:"cancel_token"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

func NewBooking(
	tenantID, serviceID, customerID, date, timeSlot string,
	duration int,
	customerName, customerPhone, customerEmail, notes string,
) *Booking {
	return &Booking{
		TenantID:       tenantID,
		ServiceID:      serviceID,
		CustomerID:     customerID,
		Date:           date,
		TimeSlot:       timeSlot,
		Duration:       duration,
		Status:         "pending",
		CustomerName:   customerName,
		CustomerPhone:  customerPhone,
		CustomerEmail:  customerEmail,
		Notes:          notes,
		CancelToken:    generateCancelToken(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func generateCancelToken() string {
	token := uuid.New().String()
	return token[:8] + token[len(token)-8:]
}