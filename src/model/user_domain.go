package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetJSONValue() (string, error)
	SetId(string)
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (ud *userDomain) SetId(id string) {
	ud.ID = id
}

// Request e Response são usados para comunicação com o mundo externo, enquanto Domain é usado para lógica de negócio interna
type userDomain struct {
	ID       string
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func (ud *userDomain) GetEmail() string {
	return ud.Email
}
func (ud *userDomain) GetPassword() string {
	return ud.Password
}
func (ud *userDomain) GetName() string {
	return ud.Name
}
func (ud *userDomain) GetAge() int8 {
	return ud.Age
}

// Implementação do método EncryptPassword para a estrutura UserDomain, que criptografa a senha usando MD5
func (ud *userDomain) EncryptPassword() {
	// Criptografa a senha usando MD5 e armazena o resultado de volta na estrutura UserDomain
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
