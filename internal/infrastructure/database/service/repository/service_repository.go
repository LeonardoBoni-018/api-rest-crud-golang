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
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository/entity"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository/entity/converter"
)

const serviceColletionEnv = "MONGODB_SERVICE_COLLECTION"

type serviceRepository struct {
	database *mongo.Database
}

func NewServiceRepository(database *mongo.Database) ServiceRepository {
	return &serviceRepository{database}
}

type ServiceRepository interface {
	CreateService(service *domain.Service) (*domain.Service, *rest_err.RestErr)
	FindServicesByTenantID(tenantID string) ([]*domain.Service, *rest_err.RestErr)
	FindServiceByID(id string) (*domain.Service, *rest_err.RestErr)
	UpdateService(id string, service *domain.Service) (*domain.Service, *rest_err.RestErr)
	DeleteService(id string) *rest_err.RestErr
}

func (r *serviceRepository) colletion() *mongo.Collection {
	colletionName := os.Getenv(serviceColletionEnv)
	return r.database.Collection(colletionName)
}

func (r *serviceRepository) CreateService(s *domain.Service) (*domain.Service, *rest_err.RestErr) {
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()

	result, err := r.colletion().InsertOne(context.Background(), converter.ToEntity(s))

	if err != nil {
		logger.Error("Error trying to insert service", err, zap.String("journey", "createService"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	s.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return s, nil
}

func (r *serviceRepository) FindServicesByTenantID(tenantID string) ([]*domain.Service, *rest_err.RestErr) {

	filter := bson.D{{Key: "tenant_id", Value: tenantID}}
	cursor, err := r.colletion().Find(context.Background(), filter)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	defer cursor.Close(context.Background())

	var services []*domain.Service

	for cursor.Next(context.Background()) {
		var entity entity.ServiceEntity
		if err := cursor.Decode(&entity); err != nil {
			return nil, rest_err.NewInternalServerError(err.Error())
		}
		services = append(services, converter.ToDomain(entity))
	}

	return services, nil
}

func (r *serviceRepository) FindServiceByID(id string) (*domain.Service, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid service id")
	}

	filter := bson.D{{Key: "_id", Value: objID}}
	doc := &entity.ServiceEntity{}

	if err := r.colletion().FindOne(context.Background(), filter).Decode(doc); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("service not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return converter.ToDomain(*doc), nil
}

func (r *serviceRepository) UpdateService(id string, s *domain.Service) (*domain.Service, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid service id")
	}

	s.UpdatedAt = time.Now()
	_, err = r.colletion().UpdateOne(
		context.Background(),
		bson.D{{Key: "_id", Value: objID}},
		bson.D{{Key: "$set", Value: converter.ToEntity(s)}},
	)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	s.ID = id
	return s, nil
}

func (r *serviceRepository) DeleteService(id string) *rest_err.RestErr {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return rest_err.NewBadRequestError("invalid service id")
	}

	_, err = r.colletion().DeleteOne(context.Background(), bson.D{{Key: "_id", Value: objID}})
	if err != nil {
		return rest_err.NewInternalServerError(err.Error())
	}
	return nil
}
