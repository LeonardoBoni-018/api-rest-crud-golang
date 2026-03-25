package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository/entity"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"

)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
