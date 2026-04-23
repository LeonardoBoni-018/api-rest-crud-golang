package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/loger"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/configuration/database/mongodb"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/routes"
)

func main() {
	loger.Info("About to start user application")
	// Carrega as variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to conect to database, error=%s /n",
			err.Error())
		return
	}

	userController, tenantController, serviceController, bookingController, dashboardController, chatController := initDependencies(database)

	router := gin.Default()
	// Inicializa as rotas
	routes.InitRoutes(&router.RouterGroup, userController, tenantController, serviceController, bookingController, dashboardController, chatController)
	// Inicializa o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error running server", err)
	}
}
