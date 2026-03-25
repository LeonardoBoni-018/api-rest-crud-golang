package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository/entity"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email,
		entity.Password,
		entity.Name,
		entity.Age,
	)

	// domain.SetID(entity.ID.Hex())
	return domain
}
