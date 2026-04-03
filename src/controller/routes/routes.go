package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model"
)

// Inicializa as rotas
func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", model.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", model.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
	r.POST("/businesses", controller.CreateBusiness)
	r.POST("/businesses/:businessId/services", controller.CreateService)
	r.GET("/businesses/:businessId/services", controller.GetServicesByBusiness)

}
