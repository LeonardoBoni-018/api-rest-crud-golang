package notification

import "time"

type Notification struct {
	ID          string    `json:"id" bson:"_id"`
	TenantID    string    `json:"tenant_id" bson:"tenant_id"`
	BookingID  string    `json:"booking_id" bson:"booking_id"`
	Type       string    `json:"type" bson:"type"` // confirmation, reminder_24h, reminder_1h, cancellation, reschedule
	Status     string    `json:"status" bson:"status"` // pending, sent, failed
	To         string    `json:"to" bson:"to"`
	ToName     string    `json:"to_name" bson:"to_name"`
	Subject    string    `json:"subject" bson:"subject"`
	Body       string    `json:"body" bson:"body"`
	SentAt    *time.Time `json:"sent_at" bson:"sent_at"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type EmailTemplate struct {
	Type    string `json:"type" bson:"type"`
	Subject string `json:"subject" bson:"subject"`
	Body    string `json:"body" bson:"body"`
}

func NewNotification(
	tenantID, bookingID, notificationType, to, toName, subject, body string,
) *Notification {
	return &Notification{
		TenantID:   tenantID,
		BookingID:  bookingID,
		Type:       notificationType,
		Status:     "pending",
		To:         to,
		ToName:     toName,
		Subject:   subject,
		Body:      body,
		CreatedAt:  time.Now(),
		UpdatedAt: time.Now(),
	}
}

const (
	TypeBookingConfirmation = "confirmation"
	TypeBookingReminder24h   = "reminder_24h"
	TypeBookingReminder1h  = "reminder_1h"
	TypeBookingCancellation  = "cancellation"
	TypeBookingReschedule  = "reschedule"
)

const (
	StatusPending = "pending"
	StatusSent   = "sent"
	StatusFailed = "failed"
)