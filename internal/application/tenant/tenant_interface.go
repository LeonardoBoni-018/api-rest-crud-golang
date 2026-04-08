package tenant

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	tenantRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
)

type TenantService interface {
	CreateTenant(*tenantDomain.Tenant) (*tenantDomain.Tenant, *rest_err.RestErr)
	FindTenantByID(string) (*tenantDomain.Tenant, *rest_err.RestErr)
	FindTenantBySlug(string) (*tenantDomain.Tenant, *rest_err.RestErr)
	UpdateTenant(string, *tenantDomain.Tenant) *rest_err.RestErr
}

type tenantService struct {
	tenantRepository tenantRepo.TenantRepository
}

func NewTenantService(repo tenantRepo.TenantRepository) TenantService {
	return &tenantService{tenantRepository: repo}
}
