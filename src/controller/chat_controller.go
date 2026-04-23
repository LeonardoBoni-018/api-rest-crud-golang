package controller

import (
	"net/http"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	chatapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/chat"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/response"
	"github.com/gin-gonic/gin"
)

type ChatControllerInterface interface {
	StartConversation(c *gin.Context)
	SendMessage(c *gin.Context)
	GetConversation(c *gin.Context)
	ListConversations(c *gin.Context)
}

type chatController struct {
	chat chatapp.ChatService
}

func NewChatController(service chatapp.ChatService) ChatControllerInterface {
	return &chatController{chat: service}
}

func (cc *chatController) StartConversation(c *gin.Context) {
	tenantSlug := c.Param("slug")
	var req request.StartChatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	conversation, restErr := cc.chat.StartConversation(tenantSlug, req.CustomerID, req.Message)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusCreated, response.ChatConversationResponse{Conversation: conversation})
}

func (cc *chatController) SendMessage(c *gin.Context) {
	tenantSlug := c.Param("slug")
	conversationID := c.Param("conversationId")
	var req request.ChatMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	if conversationID == "" {
		conversationID = req.ConversationID
	}

	if conversationID == "" {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("conversation_id is required"))
		return
	}

	conversation, restErr := cc.chat.SendMessage(tenantSlug, conversationID, req.Message)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response.ChatConversationResponse{Conversation: conversation})
}

func (cc *chatController) GetConversation(c *gin.Context) {
	tenantSlug := c.Param("slug")
	conversationID := c.Param("conversationId")
	conversation, restErr := cc.chat.GetConversation(tenantSlug, conversationID)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response.ChatConversationResponse{Conversation: conversation})
}

func (cc *chatController) ListConversations(c *gin.Context) {
	tenantSlug := c.Param("slug")
	conversations, restErr := cc.chat.ListConversations(tenantSlug)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response.ChatConversationsResponse{Conversations: conversations})
}
