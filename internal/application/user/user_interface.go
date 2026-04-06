package user

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	repository "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository"
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
	CreateUserServices(user.UserDomainInterface) *rest_err.RestErr
	UpdateUserDomain(string, user.UserDomainInterface) *rest_err.RestErr
	FindUserByIdServices(
		id string,
	) (user.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailServices(
		email string,
	) (user.UserDomainInterface, *rest_err.RestErr)
	DeleteUserDomain(string) *rest_err.RestErr
	LoginUserServices(
		userDomain user.UserDomainInterface,
	) (user.UserDomainInterface, string, *rest_err.RestErr)
}
