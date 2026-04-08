package user

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":        ud.ID,
		"email":     ud.Email,
		"name":      ud.Name,
		"age":       ud.Age,
		"tenant_id": ud.TenantID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
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
func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(tokenValue, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}
		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		errRest := rest_err.NewUnauthorizedError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorizedError("Invalid token")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	ageFloat, ok := claims["age"].(float64)
	if !ok {
		errRest := rest_err.NewUnauthorizedError("Invalid token: age claim type")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	// Extrair tenant_id
	tenantID := ""
	if tID, ok := claims["tenant_id"].(string); ok {
		tenantID = tID
	}

	fmt.Printf("DEBUG - tenant_id from JWT: '%s'\n", tenantID)
	if tenantID == "" {
		errRest := rest_err.NewBadRequestError("tenant_id not found in token - please register a tenant first")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	userDomain := &userDomain{
		ID:       claims["id"].(string),
		Email:    claims["email"].(string),
		Name:     claims["name"].(string),
		Age:      int8(ageFloat),
		TenantID: tenantID,
	}

	c.Set("user", userDomain)
	c.Set("user_id", userDomain.ID)
	c.Set("tenant_id", userDomain.TenantID)
	c.Next()
}
