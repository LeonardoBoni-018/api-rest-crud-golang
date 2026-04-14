package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingEntity struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TenantID      string             `json:"tenant_id" bson:"tenant_id,omitempty"`
	ServiceID     string             `json:"service_id" bson:"service_id,omitempty"`
	CustomerID    string             `json:"customer_id" bson:"customer_id,omitempty"`
	Date          string             `json:"date" bson:"date,omitempty"`
	TimeSlot      string             `json:"time_slot" bson:"time_slot,omitempty"`
	Duration      int                `json:"duration" bson:"duration,omitempty"`
	Status        string             `json:"status" bson:"status,omitempty"`
	CustomerName  string             `json:"customer_name" bson:"customer_name,omitempty"`
	CustomerPhone string             `json:"customer_phone" bson:"customer_phone,omitempty"`
	CustomerEmail string             `json:"customer_email" bson:"customer_email,omitempty"`
	Notes         string             `json:"notes" bson:"notes,omitempty"`
	CreatedAt     primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt     primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}
