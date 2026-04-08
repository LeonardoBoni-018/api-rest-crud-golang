package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository/entity"
)

func ConvertDomainToEntity(
	domain user.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		TenantID: domain.GetTenantID(),
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
