package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) user.UserDomainInterface {
	domain := user.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	domain.SetId(entity.ID.Hex())
	domain.SetTenantID(entity.TenantID)
	return domain
}
