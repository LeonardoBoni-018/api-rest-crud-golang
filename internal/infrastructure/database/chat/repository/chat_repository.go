package repository

import (
	"context"
	"os"
	"time"

	"github.com/bytedance/gopkg/util/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	chatdomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/chat"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/chat/repository/entity"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/chat/repository/entity/converter"
)

const chatCollectionEnv = "MONGODB_CHAT_COLLECTION"

type chatRepository struct {
	database *mongo.Database
}

type ChatRepository interface {
	CreateConversation(conversation *chatdomain.Conversation) (*chatdomain.Conversation, *rest_err.RestErr)
	FindConversationByID(id string) (*chatdomain.Conversation, *rest_err.RestErr)
	AddMessage(conversationID string, message chatdomain.Message) (*chatdomain.Conversation, *rest_err.RestErr)
	FindConversationsByTenantID(tenantID string) ([]*chatdomain.Conversation, *rest_err.RestErr)
}

func NewChatRepository(database *mongo.Database) ChatRepository {
	return &chatRepository{database: database}
}

func (r *chatRepository) collection() *mongo.Collection {
	collectionName := os.Getenv(chatCollectionEnv)
	return r.database.Collection(collectionName)
}

func (r *chatRepository) CreateConversation(conversation *chatdomain.Conversation) (*chatdomain.Conversation, *rest_err.RestErr) {
	conversation.CreatedAt = time.Now()
	conversation.UpdatedAt = time.Now()

	result, err := r.collection().InsertOne(context.Background(), converter.ToEntity(conversation))
	if err != nil {
		logger.Error("Error trying to insert conversation", err, zap.String("journey", "createConversation"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	conversation.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return conversation, nil
}

func (r *chatRepository) FindConversationByID(id string) (*chatdomain.Conversation, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid conversation id")
	}

	doc := &entity.ConversationEntity{}
	if err := r.collection().FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("conversation not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	conversation := converter.ToDomain(*doc)
	return conversation, nil
}

func (r *chatRepository) AddMessage(conversationID string, message chatdomain.Message) (*chatdomain.Conversation, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(conversationID)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid conversation id")
	}

	messageEntity := converter.ToMessageEntity(message)

	update := bson.D{
		{Key: "$push", Value: bson.M{"messages": messageEntity}},
		{Key: "$set", Value: bson.M{"updated_at": time.Now()}},
	}

	_, err = r.collection().UpdateOne(context.Background(), bson.D{{Key: "_id", Value: objID}}, update)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return r.FindConversationByID(conversationID)
}

func (r *chatRepository) FindConversationsByTenantID(tenantID string) ([]*chatdomain.Conversation, *rest_err.RestErr) {
	filter := bson.D{{Key: "tenant_id", Value: tenantID}}
	cursor, err := r.collection().Find(context.Background(), filter)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	defer cursor.Close(context.Background())

	var conversations []*chatdomain.Conversation
	for cursor.Next(context.Background()) {
		var entity entity.ConversationEntity
		if err := cursor.Decode(&entity); err != nil {
			return nil, rest_err.NewInternalServerError(err.Error())
		}
		conversations = append(conversations, converter.ToDomain(entity))
	}
	return conversations, nil
}
