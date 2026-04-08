package user

import "github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetTenantID() string
	GetJSONValue() (string, error)
	SetId(string)
	SetTenantID(string)
	EncryptPassword()
	GetId() string
	GenerateToken() (string, *rest_err.RestErr)
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

func NewUserLoginDomain(
	email, password string,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
	}
}

func NewUserUpdateDomain(
	age int8,
	name string,
) UserDomainInterface {
	return &userDomain{
		Name: name,
		Age:  age,
	}
}
