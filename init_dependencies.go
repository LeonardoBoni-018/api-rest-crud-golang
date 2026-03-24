package main

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model/service"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
