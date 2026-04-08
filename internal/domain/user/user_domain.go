package user

import (
	"encoding/json"
	"fmt"
)

// Request e Response são usados para comunicação com o mundo externo, enquanto Domain é usado para lógica de negócio interna
type userDomain struct {
	ID       string
	TenantID string
	Email    string
	Password string
	Role     string
	Name     string
	Age      int8
}

func (ud *userDomain) SetId(id string) {
	ud.ID = id
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

func (ud *userDomain) GetId() string {
	return ud.ID
}

func (ud *userDomain) GetTenantID() string {
	return ud.TenantID
}

func (ud *userDomain) SetTenantID(tenantID string) {
	ud.TenantID = tenantID
}
