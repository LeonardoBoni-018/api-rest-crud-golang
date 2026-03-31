package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetJSONValue() (string, error)
	SetId(string)
	EncryptPassword()
	GetId() string
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
