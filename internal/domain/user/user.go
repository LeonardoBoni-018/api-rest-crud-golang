package user

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
)

type User struct {
	ID        string    `json:"id" bson:"_id"`
	TenantID  string    `json:"tenant_id" bson:"tenant_id"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"-" bson:"password"`
	Role      string    `json:"role" bson:"role"`
	Name      string    `json:"name" bson:"name"`
	Age       int8      `json:"age" bson:"age"`
	Phone     string    `json:"phone" bson:"phone"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func (u *User) SetId(id string) {
	u.ID = id
}

func (u *User) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetPassword() string {
	return u.Password
}

func (u *User) GetName() string {
	return u.Name
}

func (u *User) GetAge() int8 {
	return u.Age
}

func (u *User) GetId() string {
	return u.ID
}

func (u *User) GetTenantID() string {
	return u.TenantID
}

func (u *User) SetTenantID(tenantID string) {
	u.TenantID = tenantID
}

func (u *User) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}

func (u *User) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":        u.ID,
		"email":     u.Email,
		"name":      u.Name,
		"age":       u.Age,
		"tenant_id": u.TenantID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("Error trying to generate jwt token, err =%s", err.Error()))
	}

	return tokenString, nil
}
