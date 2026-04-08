package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TenantSettingsEntity struct {
	BusinessType    string   `json:"business_type" bson:"business_type"`
	WorkingDays     []string `json:"working_days" bson:"working_days"`
	WorkingHours    string   `json:"working_hours" bson:"working_hours"`
	BookingInterval int      `json:"booking_interval" bson:"booking_interval"`
	MaxAdvanceDays  int      `json:"max_advance_days" bson:"max_advance_days"`
}

type TenantEntity struct {
	ID        primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
	Name      string               `json:"name" bson:"name"`
	Slug      string               `json:"slug" bson:"slug"`
	Email     string               `json:"email" bson:"email"`
	Phone     string               `json:"phone" bson:"phone"`
	Plan      string               `json:"plan" bson:"plan"`
	Status    string               `json:"status" bson:"status"`
	Settings  TenantSettingsEntity `json:"settings" bson:"settings"`
	CreatedAt time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time            `json:"updated_at" bson:"updated_at"`
}
