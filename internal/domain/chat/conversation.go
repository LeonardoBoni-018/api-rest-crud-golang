package chat

import "time"

type Message struct {
	ID        string    `json:"id" bson:"_id"`
	Role      string    `json:"role" bson:"role"`
	Content   string    `json:"content" bson:"content"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
}

type Conversation struct {
	ID         string    `json:"id" bson:"_id"`
	TenantID   string    `json:"tenant_id" bson:"tenant_id"`
	CustomerID string    `json:"customer_id" bson:"customer_id"`
	Messages   []Message `json:"messages" bson:"messages"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

func NewMessage(id, role, content string) Message {
	return Message{
		ID:        id,
		Role:      role,
		Content:   content,
		Timestamp: time.Now(),
	}
}

func NewConversation(tenantID, customerID string, messages []Message) *Conversation {
	return &Conversation{
		TenantID:   tenantID,
		CustomerID: customerID,
		Messages:   messages,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
