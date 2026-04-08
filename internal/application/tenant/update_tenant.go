package tenant

import (
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
)

func (s *tenantService) UpdateTenant(id string, tenant *tenantDomain.Tenant) *rest_err.RestErr {
	tenant.UpdatedAt = time.Now()
	_, err := s.tenantRepository.UpdateTenant(id, tenant)
	return err
}
