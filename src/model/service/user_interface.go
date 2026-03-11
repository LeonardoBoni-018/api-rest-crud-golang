package service

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUserDomain(model.UserDomainInterface) *rest_err.RestErr
	UpdateUserDomain(string, model.UserDomainInterface) *rest_err.RestErr
	FindUserDomain(string) (*model.UserDomainInterface, *rest_err.RestErr)
	DeleteUserDomain(string) *rest_err.RestErr
}
