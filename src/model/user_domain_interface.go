package model

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
