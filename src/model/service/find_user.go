package service

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

// func (*userDomainService) FindUserDomain(string) (*model.UserDomainInterface, *rest_err.RestErr) {
// 	return nil, nil
// }

func (ud *userDomainService) FindUserByIdServices(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserById services", zap.String("journey", "findUserById"))
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindUserByEmail services", zap.String("journey", "findUserByEmail"))
	return ud.userRepository.FindUserByEmail(email)
}
