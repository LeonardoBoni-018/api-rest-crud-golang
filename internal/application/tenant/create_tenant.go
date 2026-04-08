package tenant

import (
	"strings"
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
)

func (s *tenantService) CreateTenant(t *tenantDomain.Tenant) (*tenantDomain.Tenant, *rest_err.RestErr) {
	if t.Name == "" || t.Slug == "" || t.Email == "" {
		return nil, rest_err.NewBadRequestError("name, slug and email are required")
	}

	t.Slug = strings.ToLower(strings.TrimSpace(t.Slug))
	existing, err := s.tenantRepository.FindTenantBySlug(t.Slug)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, rest_err.NewBadRequestError("tenant with this slug already exists")
	}

	t.Status = "active"
	if t.Plan == "" {
		t.Plan = "free"
	}

	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	return s.tenantRepository.CreateTenant(t)
}
