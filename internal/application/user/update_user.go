package user

import (
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
)

func (ud *userDomainService) UpdateUserDomain(userId string, userDomain user.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init UpdateUserDomain", zap.String("journey", "updateUser"))

	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Error trying to call repository", err, zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("User Updated successfully", zap.String("journey", "updateUser"))
	return nil
}
