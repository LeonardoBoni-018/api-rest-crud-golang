package main

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/model/service"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserDomainService(userRepo)

	businessRepo := repository.NewBusinessRepository(database)
	businessService := service.NewBusinessService(businessRepo)
	controller.SetBusinessService(businessService)

	serviceRepo := repository.NewServiceRepository(database)
	serviceService := service.NewServiceService(serviceRepo)
	controller.SetServiceService(serviceService)

	return controller.NewUserControllerInterface(userService)
}
