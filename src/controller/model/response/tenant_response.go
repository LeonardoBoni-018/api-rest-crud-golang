package response

import (
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	userDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
)

type TenantResponse struct {
	Tenant *tenantDomain.Tenant `json:"tenant"`
}

type TenantOnboardResponse struct {
	Tenant *tenantDomain.Tenant `json:"tenant"`
	Owner  *userDomain.User     `json:"owner"`
	Token  string               `json:"token"`
}
