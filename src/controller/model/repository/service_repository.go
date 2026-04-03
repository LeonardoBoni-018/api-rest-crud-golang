package repository

import (
	"context"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceRepository struct {
	DB *mongo.Database
}

func NewServiceRepository(db *mongo.Database) *ServiceRepository {
	return &ServiceRepository{DB: db}
}

func (r *ServiceRepository) Create(ctx context.Context, service model.ServiceDomain) (*mongo.InsertOneResult, error) {
	return r.DB.Collection("services").InsertOne(ctx, service)
}

func (r *ServiceRepository) FindByBusinessID(ctx context.Context, businessID string) ([]model.ServiceDomain, error) {
	cursor, err := r.DB.Collection("services").Find(ctx, bson.M{"business_id": businessID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	services := make([]model.ServiceDomain, 0)
	for cursor.Next(ctx) {
		var item model.ServiceDomain
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		services = append(services, item)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return services, nil
}
