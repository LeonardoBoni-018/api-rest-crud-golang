package service

import "time"

type Service struct {
	ID          string    `json:"id" bson:"_id"`
	TenantID    string    `json:"tenant_id" bson:"tenant_id"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Duration    int       `json:"duration" bson:"duration"`
	Price       float64   `json:"price" bson:"price"`
	IsActive    bool      `json:"is_active" bson:"is_active"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func NewService(
	tenantID, name, description string,
	duration int,
	price float64,
) *Service {
	return &Service{
		TenantID:    tenantID,
		Name:        name,
		Description: description,
		Duration:    duration,
		Price:       price,
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
