package service

import (
	"context"
	"errors"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceService struct {
	Repo *repository.ServiceRepository
}

func NewServiceService(repo *repository.ServiceRepository) *ServiceService {
	return &ServiceService{Repo: repo}
}

func (s *ServiceService) Create(ctx context.Context, service model.ServiceDomain) (*mongo.InsertOneResult, error) {
	if service.Name == "" {
		return nil, errors.New("service name required")
	}
	if service.BusinessID == "" {
		return nil, errors.New("business id required")
	}
	if service.DurationMin <= 0 {
		return nil, errors.New("duration must be positive")
	}
	if service.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}

	service.CreatedAt = service.CreatedAt.UTC()
	service.UpdatedAt = service.CreatedAt

	return s.Repo.Create(ctx, service)
}

func (s *ServiceService) FindByBusinessID(ctx context.Context, businessID string) ([]model.ServiceDomain, error) {
	if businessID == "" {
		return nil, errors.New("business id required")
	}
	return s.Repo.FindByBusinessID(ctx, businessID)
}
