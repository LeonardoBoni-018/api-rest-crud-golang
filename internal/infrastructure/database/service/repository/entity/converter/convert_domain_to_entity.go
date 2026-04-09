package converter

import (
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository/entity"

)

func ToEntity(domain *service.Service) *entity.ServiceEntity {
	entity := &entity.ServiceEntity{
		TenantID:    domain.TenantID,
		Name:        domain.Name,
		Description: domain.Description,
		Duration:    domain.Duration,
		Price:       domain.Price,
		IsActive:    domain.IsActive,
		CreatedAt:   primitive.NewDateTimeFromTime(domain.CreatedAt),
		UpdatedAt:   primitive.NewDateTimeFromTime(domain.UpdatedAt),
	}

	if domain.ID != "" {
		objectID, _ := primitive.ObjectIDFromHex(domain.ID)
		entity.ID = objectID
	}
	return entity
}
