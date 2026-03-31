package service

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

func (ud *userDomainService) LoginUserServices(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init LoginUserServices", zap.String("journey", "loginUser"))

	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, err
	}

	logger.Info("LoginUser service executed successfully", zap.String("userId", user.GetId()), zap.String("journey", "loginUser"))
	return user, nil
}
