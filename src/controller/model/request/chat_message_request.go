package request

type ChatMessageRequest struct {
	ConversationID string `json:"conversation_id"`
	Message        string `json:"message" binding:"required"`
}
