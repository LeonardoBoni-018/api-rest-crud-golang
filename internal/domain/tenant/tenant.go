package tenant

import "time"

// internal/domain/tenant/tenant.go
type Tenant struct {
	ID        string    `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Slug      string    `json:"slug" bson:"slug"` // barbearia-joao
	Email     string    `json:"email" bson:"email"`
	Phone     string    `json:"phone" bson:"phone"`
	Plan      string    `json:"plan" bson:"plan"`     // free, basic, pro
	Status    string    `json:"status" bson:"status"` // active, suspended
	Settings  Settings  `json:"settings" bson:"settings"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Settings struct {
	BusinessType    string   `json:"business_type"`    // salon, barbershop, clinic
	WorkingDays     []string `json:"working_days"`     // ["mon", "tue", "wed"]
	WorkingHours    string   `json:"working_hours"`    // "09:00-18:00"
	BookingInterval int      `json:"booking_interval"` // 30 minutos
	MaxAdvanceDays  int      `json:"max_advance_days"` // 30 dias
}

