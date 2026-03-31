package model

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"

)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.ID,
		"email": ud.Email,
		"name":  ud.Name,
		"age":   ud.Age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expira em 24 horas
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", rest_err.NewInternalServerError(fmt.Sprintf("Error trying to generate jwt token, err =%s", err.Error()))
	}
	return tokenString, nil
}
