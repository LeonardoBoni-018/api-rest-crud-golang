package service

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"

)

func (ud *userDomainService) CreateUserDomain(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init CreateUserDomain", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	_, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		return err
	}

	return nil
}