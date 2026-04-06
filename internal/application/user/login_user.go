package user

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
)

func (ud *userDomainService) LoginUserServices(
	userDomain user.UserDomainInterface,
) (user.UserDomainInterface, string, *rest_err.RestErr) {
	logger.Info("Init LoginUserServices", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}

	logger.Info("LoginUser service executed successfully", zap.String("userId", user.GetId()), zap.String("journey", "loginUser"))
	return user, token, nil
}
