package service

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
)

func (s *serviceService) CreateService(serviceDomain *domain.Service) (*domain.Service, *rest_err.RestErr) {
	return s.serviceRepository.CreateService(serviceDomain)
}
