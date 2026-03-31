package model

import (
	"fmt"
	"os"
	"strings"
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

func VerifyToken(tokenValue string) (UserDomainInterface, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	tokenValue = RemoveBearerPrefix(tokenValue)

	token, err := jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, rest_err.NewUnauthorizedError("Invalid token")
	}

	// idRaw, exists := claims["id"]
	// if !exists {
	// 	return nil, rest_err.NewUnauthorizedError("Invalid token: id missing")
	// }

	// var id string
	// switch v := idRaw.(type) {
	// case string:
	// 	id = v
	// case float64:
	// 	id = fmt.Sprintf("%.0f", v)
	// default:
	// 	return nil, rest_err.NewUnauthorizedError("Invalid token: id claim type")
	// }

	ageRaw, ageExists := claims["age"]
	if !ageExists {
		return nil, rest_err.NewUnauthorizedError("Invalid token: age missing")
	}

	ageFloat, ok := ageRaw.(float64)
	if !ok {
		return nil, rest_err.NewUnauthorizedError("Invalid token: age claim type")
	}

	return &userDomain{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
		Name:  claims["name"].(string),
		Age:   int8(ageFloat),
	}, nil
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer") {
		token = strings.TrimPrefix(token, "Bearer ")
	}
	return token
}
