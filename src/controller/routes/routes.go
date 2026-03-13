package routes

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
	"github.com/gin-gonic/gin"
)

// Inicializa as rotas
func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {

	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
