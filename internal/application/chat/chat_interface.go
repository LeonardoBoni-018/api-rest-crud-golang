package chat

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	chatdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/chat"
	aiClient "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/ai"
	chatRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/chat/repository"
	tenantRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
)

type ChatService interface {
	StartConversation(tenantSlug, customerID, message string) (*chatdomain.Conversation, *rest_err.RestErr)
	SendMessage(tenantSlug, conversationID, message string) (*chatdomain.Conversation, *rest_err.RestErr)
	GetConversation(tenantSlug, conversationID string) (*chatdomain.Conversation, *rest_err.RestErr)
	ListConversations(tenantSlug string) ([]*chatdomain.Conversation, *rest_err.RestErr)
}

type chatService struct {
	chatRepository   chatRepo.ChatRepository
	tenantRepository tenantRepo.TenantRepository
	openAIClient     *aiClient.OpenAIClient
}

func NewChatService(
	chatRepository chatRepo.ChatRepository,
	tenantRepository tenantRepo.TenantRepository,
	openAIClient *aiClient.OpenAIClient,
) ChatService {
	return &chatService{
		chatRepository:   chatRepository,
		tenantRepository: tenantRepository,
		openAIClient:     openAIClient,
	}
}
