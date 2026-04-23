package converter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	chatdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/chat"
	entity "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/chat/repository/entity"
)

func ToEntity(conv *chatdomain.Conversation) entity.ConversationEntity {
	messages := make([]entity.MessageEntity, 0, len(conv.Messages))
	for _, msg := range conv.Messages {
		msgID, err := primitive.ObjectIDFromHex(msg.ID)
		if err != nil {
			msgID = primitive.NewObjectID()
		}
		messages = append(messages, entity.MessageEntity{
			ID:        msgID,
			Role:      msg.Role,
			Content:   msg.Content,
			Timestamp: msg.Timestamp,
		})
	}

	convID, err := primitive.ObjectIDFromHex(conv.ID)
	if err != nil {
		convID = primitive.NewObjectID()
	}

	return entity.ConversationEntity{
		ID:         convID,
		TenantID:   conv.TenantID,
		CustomerID: conv.CustomerID,
		Messages:   messages,
		CreatedAt:  conv.CreatedAt,
		UpdatedAt:  conv.UpdatedAt,
	}
}

func ToDomain(entityConv entity.ConversationEntity) *chatdomain.Conversation {
	messages := make([]chatdomain.Message, 0, len(entityConv.Messages))
	for _, msg := range entityConv.Messages {
		messages = append(messages, chatdomain.Message{
			ID:        msg.ID.Hex(),
			Role:      msg.Role,
			Content:   msg.Content,
			Timestamp: msg.Timestamp,
		})
	}

	return &chatdomain.Conversation{
		ID:         entityConv.ID.Hex(),
		TenantID:   entityConv.TenantID,
		CustomerID: entityConv.CustomerID,
		Messages:   messages,
		CreatedAt:  entityConv.CreatedAt,
		UpdatedAt:  entityConv.UpdatedAt,
	}
}

func ToMessageEntity(msg chatdomain.Message) entity.MessageEntity {
	msgID, _ := primitive.ObjectIDFromHex(msg.ID)
	return entity.MessageEntity{
		ID:        msgID,
		Role:      msg.Role,
		Content:   msg.Content,
		Timestamp: msg.Timestamp,
	}
}
