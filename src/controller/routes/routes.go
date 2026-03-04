package routes

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
	"github.com/gin-gonic/gin"
)

// Inicializa as rotas
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}
