package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ServiceEntity struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TenantID    string             `json:"tenant_id" bson:"tenant_id,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description" bson:"description,omitempty"`
	Duration    int                `json:"duration" bson:"duration,omitempty"`
	Price       float64            `json:"price" bson:"price,omitempty"`
	IsActive    bool               `json:"is_active" bson:"is_active,omitempty"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   primitive.DateTime `json:"updated_at" bson:"updated_at,omitempty"`
}
