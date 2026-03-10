package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

// Request e Response são usados para comunicação com o mundo externo, enquanto Domain é usado para lógica de negócio interna
type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

// Implementação do método EncryptPassword para a estrutura UserDomain, que criptografa a senha usando MD5
func (ud *UserDomain) EncryptPassword() {
	// Criptografa a senha usando MD5 e armazena o resultado de volta na estrutura UserDomain
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

// Interface para o UserDomain, definindo os métodos que devem ser implementados
type UserDomainInterface interface {
	// CRUD - Create, Read, Update, Delete
	// Cada método retorna um ponteiro para rest_err.RestErr, que é uma estrutura personalizada para lidar com erros de forma consistente em toda a aplicação
	// Os métodos de leitura normalmente devolvem o próprio domínio para que o chamador tenha acesso aos dados encontrados.
	CreateUserDomain() *rest_err.RestErr
	UpdateUserDomain(string) *rest_err.RestErr
	FindUserDomain(string) (*UserDomain, *rest_err.RestErr)
	DeleteUserDomain(string) *rest_err.RestErr
}
