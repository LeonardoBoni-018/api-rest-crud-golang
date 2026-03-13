package main

import (
	"log"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/routes"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loger.Info("About to start user application")
	// Carrega as variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Inir dependency injection
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	// Inicializa as rotas
	routes.InitRoutes(&router.RouterGroup, userController)
	// Inicializa o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error running server", err)
	}
}
