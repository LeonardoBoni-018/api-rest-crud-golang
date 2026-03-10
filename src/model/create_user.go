package model

import (
	"fmt"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/bytedance/gopkg/util/logger"
	"go.uber.org/zap"
)

// Implementação do método CreateUserDomain para a estrutura UserDomain
func (ud *UserDomain) CreateUserDomain() *rest_err.RestErr {

	// Aqui você pode adicionar a lógica para criar um usuário, como validação de dados, interação com o banco de dados, etc.
	// Validar email, confirmação em duas etapas, etc.

	logger.Info("Init CreateUserModel", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	fmt.Println()
	return nil
}
