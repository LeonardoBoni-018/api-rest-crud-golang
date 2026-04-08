package repository

import (
	"context"
	"os"
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository/converter"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const tenantCollectionEnv = "MONGODB_TENANT_COLLECTION"

type tenantRepository struct {
	database *mongo.Database
}

func NewTenantRepository(database *mongo.Database) *tenantRepository {
	return &tenantRepository{database: database}
}

type TenantRepository interface {
	CreateTenant(tenant *tenant.Tenant) (*tenant.Tenant, *rest_err.RestErr)
	FindTenantById(id string) (*tenant.Tenant, *rest_err.RestErr)
	FindTenantBySlug(slug string) (*tenant.Tenant, *rest_err.RestErr)
	UpdateTenant(id string, tenant *tenant.Tenant) (*tenant.Tenant, *rest_err.RestErr)
}

func (r *tenantRepository) colletion() *mongo.Collection {
	collectionName := os.Getenv(tenantCollectionEnv)
	return r.database.Collection(collectionName)
}

func (r *tenantRepository) CreateTenant(t *tenant.Tenant) (*tenant.Tenant, *rest_err.RestErr) {
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	result, err := r.colletion().InsertOne(context.Background(), converter.ToEntity(t))
	if err != nil {
		return nil, rest_err.NewInternalServerError("error creating tenant")
	}

	t.ID = result.InsertedID.(primitive.ObjectID).Hex()
	return t, nil
}

func (r *tenantRepository) FindTenantById(id string) (*tenant.Tenant, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid tenant ID")
	}

	doc := &entity.TenantEntity{}
	err = r.colletion().FindOne(context.Background(), bson.M{"_id": objID}).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("tenant not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return converter.ToDomain(doc), nil
}

func (r *tenantRepository) FindTenantBySlug(slug string) (*tenant.Tenant, *rest_err.RestErr) {
	doc := &entity.TenantEntity{}
	err := r.colletion().FindOne(context.Background(), bson.M{"slug": slug}).Decode(doc)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, rest_err.NewNotFoundError("tenant not found")
		}
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return converter.ToDomain(doc), nil
}

func (r *tenantRepository) UpdateTenant(id string, t *tenant.Tenant) (*tenant.Tenant, *rest_err.RestErr) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid tenant ID")
	}

	t.UpdatedAt = time.Now()
	_, err = r.colletion().UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": converter.ToEntity(t)},
	)
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	return nil, nil
}
