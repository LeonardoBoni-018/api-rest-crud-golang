package view

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/response"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
