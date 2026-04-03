package repository

import (
	"context"
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BusinessRepository struct {
	DB *mongo.Database
}

func NewBusinessRepository(db *mongo.Database) *BusinessRepository {
	return &BusinessRepository{DB: db}
}

func (r *BusinessRepository) Create(ctx context.Context, b model.BusinessDomain) (*mongo.InsertOneResult, error) {
	b.CreatedAt = time.Now().UTC()
	b.UpdatedAt = time.Now().UTC()
	return r.DB.Collection("businesses").InsertOne(ctx, b)
}

func (r *BusinessRepository) FindByID(ctx context.Context, id string) (model.BusinessDomain, error) {
	var business model.BusinessDomain
	err := r.DB.Collection("businesses").FindOne(ctx, bson.M{"_id": id}).Decode(&business)
	return business, err
}

// etc...
