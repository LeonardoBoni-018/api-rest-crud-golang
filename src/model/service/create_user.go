package service

import (
	"fmt"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserDomain(userDomain model.UserDomainInterface) *rest_err.RestErr {
	logger.Info("Init CreateUserModel", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
