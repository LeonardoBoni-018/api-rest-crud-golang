package response

import (
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
)

type PublicTenantInfo struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Slug         string   `json:"slug"`
	Phone        string   `json:"phone"`
	Email       string   `json:"email"`
	BusinessType string `json:"business_type"`
	WorkingDays  []string `json:"working_days"`
	WorkingHours string  `json:"working_hours"`
	BookingInterval int    `json:"booking_interval"`
	MaxAdvanceDays int     `json:"max_advance_days"`
}

func ConvertTenantToPublicInfo(tenant *tenantDomain.Tenant) *PublicTenantInfo {
	if tenant == nil {
		return nil
	}
	return &PublicTenantInfo{
		ID:             tenant.ID,
		Name:           tenant.Name,
		Slug:           tenant.Slug,
		Phone:          tenant.Phone,
		Email:          tenant.Email,
		BusinessType:  tenant.Settings.BusinessType,
		WorkingDays:   tenant.Settings.WorkingDays,
		WorkingHours:  tenant.Settings.WorkingHours,
		BookingInterval: tenant.Settings.BookingInterval,
		MaxAdvanceDays:  tenant.Settings.MaxAdvanceDays,
	}
}