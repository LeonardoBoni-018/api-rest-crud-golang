package main

import (
	"go.mongodb.org/mongo-driver/mongo"

	userapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/user"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := userapp.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
