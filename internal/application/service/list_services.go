package service

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
)

func (s *serviceService) FindServicesByTenantID(tenantID string) ([]*domain.Service, *rest_err.RestErr) {
	return s.serviceRepository.FindServicesByTenantID(tenantID)
}

func (s *serviceService) FindServiceByID(id string) (*domain.Service, *rest_err.RestErr) {
	return s.serviceRepository.FindServiceByID(id)
}

func (s *serviceService) UpdateService(id string, serviceDomain *domain.Service) (*domain.Service, *rest_err.RestErr) {
	return s.serviceRepository.UpdateService(id, serviceDomain)
}

func (s *serviceService) DeleteService(id string) *rest_err.RestErr {
	return s.serviceRepository.DeleteService(id)
}
