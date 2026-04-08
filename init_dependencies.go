package main

import (
	"go.mongodb.org/mongo-driver/mongo"

	tenantapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/tenant"
	userapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/user"
	tenantrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
	userrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
)

func initDependencies(database *mongo.Database) (controller.UserControllerInterface, controller.TenantControllerInterface) {
	tenantRepo := tenantrepo.NewTenantRepository(database)
	userRepo := userrepo.NewUserRepository(database)

	userService := userapp.NewUserDomainService(userRepo)
	tenantService := tenantapp.NewTenantService(tenantRepo)
	tenantOnboarding := tenantapp.NewTenantOnboardingService(tenantRepo, userRepo)

	tenantController := controller.NewTenantController(tenantService, tenantOnboarding)

	return controller.NewUserControllerInterface(userService), tenantController
}
