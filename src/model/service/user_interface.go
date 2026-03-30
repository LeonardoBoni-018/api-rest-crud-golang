package service

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"

)

func NewUserDomainService(
	userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) *rest_err.RestErr
	UpdateUserDomain(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserByIdServices(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserDomain(string) *rest_err.RestErr
}
