package main

import (
	"log"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega as variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	// Inicializa as rotas
	routes.InitRoutes(&router.RouterGroup)
	// Inicializa o servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Error running server", err)
	}
}
