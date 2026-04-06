package middleware

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	"github.com/gin-gonic/gin"
)

type tenantUserRepository interface {
	FindByID(userID string) (*user.User, error)
}

var userRepo tenantUserRepository

func TenantIsolation() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("user_id")

		// Buscar tenant do usuário
		user, err := userRepo.FindByID(userID)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		// Injetar tenant_id no contexto
		c.Set("tenant_id", user.TenantID)
		c.Next()
	}
}
