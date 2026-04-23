package chat

import (
	"context"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	chatdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/chat"
	tenantdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	aiClient "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/ai"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *chatService) StartConversation(tenantSlug, customerID, message string) (*chatdomain.Conversation, *rest_err.RestErr) {
	tenant, err := s.tenantRepository.FindTenantBySlug(tenantSlug)
	if err != nil {
		return nil, err
	}

	userMessage := chatdomain.NewMessage(primitive.NewObjectID().Hex(), "user", message)
	conversation := chatdomain.NewConversation(tenant.ID, customerID, []chatdomain.Message{userMessage})

	conversation, err = s.chatRepository.CreateConversation(conversation)
	if err != nil {
		return nil, err
	}

	assistantResponse, errAI := s.openAIClient.Chat(context.Background(), buildOpenAIMessages(tenant, conversation.Messages))
	if errAI != nil {
		return nil, rest_err.NewInternalServerError(errAI.Error())
	}

	assistantMessage := chatdomain.NewMessage(primitive.NewObjectID().Hex(), "assistant", assistantResponse)
	conversation, err = s.chatRepository.AddMessage(conversation.ID, assistantMessage)
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (s *chatService) SendMessage(tenantSlug, conversationID, message string) (*chatdomain.Conversation, *rest_err.RestErr) {
	conversation, err := s.chatRepository.FindConversationByID(conversationID)
	if err != nil {
		return nil, err
	}

	tenant, err := s.tenantRepository.FindTenantBySlug(tenantSlug)
	if err != nil {
		return nil, err
	}
	if conversation.TenantID != tenant.ID {
		return nil, rest_err.NewForbiddenError("conversation does not belong to tenant")
	}

	userMessage := chatdomain.NewMessage(primitive.NewObjectID().Hex(), "user", message)
	conversation, err = s.chatRepository.AddMessage(conversation.ID, userMessage)
	if err != nil {
		return nil, err
	}

	assistantResponse, errAI := s.openAIClient.Chat(context.Background(), buildOpenAIMessages(tenant, conversation.Messages))
	if errAI != nil {
		return nil, rest_err.NewInternalServerError(errAI.Error())
	}

	assistantMessage := chatdomain.NewMessage(primitive.NewObjectID().Hex(), "assistant", assistantResponse)
	conversation, err = s.chatRepository.AddMessage(conversation.ID, assistantMessage)
	if err != nil {
		return nil, err
	}

	return conversation, nil
}

func (s *chatService) GetConversation(tenantSlug, conversationID string) (*chatdomain.Conversation, *rest_err.RestErr) {
	conversation, err := s.chatRepository.FindConversationByID(conversationID)
	if err != nil {
		return nil, err
	}

	tenant, err := s.tenantRepository.FindTenantBySlug(tenantSlug)
	if err != nil {
		return nil, err
	}

	if conversation.TenantID != tenant.ID {
		return nil, rest_err.NewForbiddenError("conversation does not belong to tenant")
	}

	return conversation, nil
}

func (s *chatService) ListConversations(tenantSlug string) ([]*chatdomain.Conversation, *rest_err.RestErr) {
	tenant, err := s.tenantRepository.FindTenantBySlug(tenantSlug)
	if err != nil {
		return nil, err
	}

	return s.chatRepository.FindConversationsByTenantID(tenant.ID)
}

func buildOpenAIMessages(tenant *tenantdomain.Tenant, messages []chatdomain.Message) []aiClient.MessageBody {
	openAIMessages := make([]aiClient.MessageBody, 0, len(messages)+1)

	if tenant != nil {
		systemPrompt := aiClient.MessageBody{
			Role:    "system",
			Content: buildSystemPrompt(tenant),
		}
		openAIMessages = append(openAIMessages, systemPrompt)
	}

	for _, message := range messages {
		openAIMessages = append(openAIMessages, aiClient.MessageBody{
			Role:    message.Role,
			Content: message.Content,
		})
	}

	return openAIMessages
}

func buildSystemPrompt(tenant *tenantdomain.Tenant) string {
	return "Você é um assistente virtual para o negócio " + tenant.Name + ". " +
		"Responda em linguagem natural, sendo cordial e profissional. " +
		"O cliente quer saber sobre serviços, horários e agendamento. " +
		"Use as informações de contexto para o negócio e, se possível, ofereça ajuda com agendamento."
}
