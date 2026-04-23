package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageEntity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Role      string             `bson:"role" json:"role"`
	Content   string             `bson:"content" json:"content"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}

type ConversationEntity struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TenantID   string             `bson:"tenant_id" json:"tenant_id"`
	CustomerID string             `bson:"customer_id" json:"customer_id"`
	Messages   []MessageEntity    `bson:"messages" json:"messages"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at"`
}
