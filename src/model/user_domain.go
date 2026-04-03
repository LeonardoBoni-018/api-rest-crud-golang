package model

import (
	"encoding/json"
	"fmt"
	"time"
)

// Request e Response são usados para comunicação com o mundo externo, enquanto Domain é usado para lógica de negócio interna
type UserDomain struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	Name      string    `bson:"name" json:"name"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password" json:"password"`
	Age       int8      `bson:"age" json:"age"`
	Role      string    `bson:"role" json:"role"` // customer | business_owner
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type userDomain = UserDomain

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
