package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToEntity(t *tenant.Tenant) *entity.TenantEntity {
	id, _ := primitive.ObjectIDFromHex(t.ID)
	return &entity.TenantEntity{
		ID:     id,
		Name:   t.Name,
		Slug:   t.Slug,
		Email:  t.Email,
		Phone:  t.Phone,
		Plan:   t.Plan,
		Status: t.Status,
		Settings: entity.TenantSettingsEntity{
			BusinessType:    t.Settings.BusinessType,
			WorkingDays:     t.Settings.WorkingDays,
			WorkingHours:    t.Settings.WorkingHours,
			BookingInterval: t.Settings.BookingInterval,
			MaxAdvanceDays:  t.Settings.MaxAdvanceDays,
		},
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func ToDomain(e *entity.TenantEntity) *tenant.Tenant {
	return &tenant.Tenant{
		ID:     e.ID.Hex(),
		Name:   e.Name,
		Slug:   e.Slug,
		Email:  e.Email,
		Phone:  e.Phone,
		Plan:   e.Plan,
		Status: e.Status,
		Settings: tenant.Settings{
			BusinessType:    e.Settings.BusinessType,
			WorkingDays:     e.Settings.WorkingDays,
			WorkingHours:    e.Settings.WorkingHours,
			BookingInterval: e.Settings.BookingInterval,
			MaxAdvanceDays:  e.Settings.MaxAdvanceDays,
		},
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}
