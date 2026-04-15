package main

import (
	"go.mongodb.org/mongo-driver/mongo"

	bookingapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/booking"
	serviceapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/service"
	tenantapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/tenant"
	userapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/user"
	bookingrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository"
	servicerepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository"
	tenantrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
	userrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
)

func initDependencies(database *mongo.Database) (
	controller.UserControllerInterface,
	controller.TenantControllerInterface,
	controller.ServiceControllerInterface,
	controller.BookingControllerInterface,
) {
	tenantRepo := tenantrepo.NewTenantRepository(database)
	userRepo := userrepo.NewUserRepository(database)
	serviceRepo := servicerepo.NewServiceRepository(database)
	bookingRepo := bookingrepo.NewBookingRepository(database)

	userService := userapp.NewUserDomainService(userRepo)
	tenantService := tenantapp.NewTenantService(tenantRepo)
	tenantOnboarding := tenantapp.NewTenantOnboardingService(tenantRepo, userRepo)
	serviceService := serviceapp.NewServiceService(serviceRepo)
	bookingService := bookingapp.NewBookingService(bookingRepo, tenantRepo, serviceRepo)

	return controller.NewUserControllerInterface(userService),
		controller.NewTenantController(tenantService, tenantOnboarding),
		controller.NewServiceController(serviceService),
		controller.NewBookingController(bookingService)
}
