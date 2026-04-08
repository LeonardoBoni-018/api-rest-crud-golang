package tenant

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
)

func (s *tenantService) FindTenantByID(id string) (*tenantDomain.Tenant, *rest_err.RestErr) {
	return s.tenantRepository.FindTenantById(id)
}

func (s *tenantService) FindTenantBySlug(slug string) (*tenantDomain.Tenant, *rest_err.RestErr) {
	return s.tenantRepository.FindTenantBySlug(slug)
}
