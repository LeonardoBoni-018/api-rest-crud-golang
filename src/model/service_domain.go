package model

import "time"

type ServiceDomain struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	BusinessID  string    `bson:"business_id" json:"business_id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	DurationMin int       `bson:"duration_min" json:"duration_min"` // em minutos
	Price       float64   `bson:"price" json:"price"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
