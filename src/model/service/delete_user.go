package service

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
)

func (ud *userDomainService) DeleteUserDomain(userId string) *rest_err.RestErr {
	logger.Info("Init DeleteUserDomain", zap.String("journey", "deleteUser"))

	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("User Deleted successfully", zap.String("journey", "deleteUser"))
	return nil
}
