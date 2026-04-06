package user

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
)

func (ud *userDomainService) CreateUserServices(
	userDomain user.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUserDomain", zap.String("journey", "createUser"))

	user, _ := ud.FindUserByEmailServices(userDomain.GetEmail())

	if user != nil {
		return rest_err.NewBadRequestError("email already exists")
	}

	userDomain.EncryptPassword()

	_, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return err
	}

	return nil
}
