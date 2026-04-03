package model

import "time"

type AppointmentDomain struct {
	ID         string    `bson:"_id,omitempty" json:"id"`
	BusinessID string    `bson:"business_id" json:"business_id"`
	ServiceID  string    `bson:"service_id" json:"service_id"`
	ClientID   string    `bson:"client_id" json:"client_id"`
	StartsAt   time.Time `bson:"starts_at" json:"starts_at"`
	EndsAt     time.Time `bson:"ends_at" json:"ends_at"`
	Status     string    `bson:"status" json:"status"` // booked, cancelled, done
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}
