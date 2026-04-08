package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
)

// Inicializa as rotas
func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
	tenantController controller.TenantControllerInterface,
) {

	r.GET("/getUserById/:userId", user.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", user.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
	r.POST("/tenants/register", tenantController.RegisterTenant)
	r.GET("/tenants/me", user.VerifyTokenMiddleware, tenantController.GetTenantMe)
	r.PATCH("/tenants/settings", user.VerifyTokenMiddleware, tenantController.UpdateTenantSettings)
}
