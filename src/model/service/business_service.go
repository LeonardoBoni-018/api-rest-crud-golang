package service

import (
	"context"
	"errors"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type BusinessService struct {
	Repo *repository.BusinessRepository
}

func NewBusinessService(repo *repository.BusinessRepository) *BusinessService {
	return &BusinessService{Repo: repo}
}

func (s *BusinessService) Create(ctx context.Context, b model.BusinessDomain) (*mongo.InsertOneResult, error) {
	if b.Name == "" {
		return nil, errors.New("name required")
	}
	return s.Repo.Create(ctx, b)
}
