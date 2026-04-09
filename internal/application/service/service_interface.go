package service

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
	repository "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository"
)

type ServiceService interface {
	CreateService(*domain.Service) (*domain.Service, *rest_err.RestErr)
	FindServicesByTenantID(string) ([]*domain.Service, *rest_err.RestErr)
	FindServiceByID(string) (*domain.Service, *rest_err.RestErr)
	UpdateService(string, *domain.Service) (*domain.Service, *rest_err.RestErr)
	DeleteService(string) *rest_err.RestErr
}

type serviceService struct {
	serviceRepository repository.ServiceRepository
}

func NewServiceService(repo repository.ServiceRepository) ServiceService {
	return &serviceService{serviceRepository: repo}
}
