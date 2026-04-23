package response

import chatdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/chat"

type ChatConversationResponse struct {
	Conversation *chatdomain.Conversation `json:"conversation"`
}

type ChatConversationsResponse struct {
	Conversations []*chatdomain.Conversation `json:"conversations"`
}
