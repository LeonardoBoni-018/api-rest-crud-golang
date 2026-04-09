package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository/entity"

)

func ToDomain(entity entity.ServiceEntity) *service.Service {
    return &service.Service{
        ID:          entity.ID.Hex(),
        TenantID:    entity.TenantID,
        Name:        entity.Name,
        Description: entity.Description,
        Duration:    entity.Duration,
        Price:       entity.Price,
        IsActive:    entity.IsActive,
        CreatedAt:   entity.CreatedAt.Time(),
        UpdatedAt:   entity.UpdatedAt.Time(),
    }
}
