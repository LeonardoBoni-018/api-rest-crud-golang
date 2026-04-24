package main

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"

	bookingapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/booking"
	chatapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/chat"
	dashboardapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/dashboard"
	serviceapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/service"
	tenantapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/tenant"
	userapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/user"
	paymentapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/payment"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/ai"
	bookingrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/booking/repository"
	chatrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/chat/repository"
	servicerepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/service/repository"
	tenantrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
	userrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository"
	paymentrepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/payment/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/payment"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
)

func initDependencies(database *mongo.Database) (
	controller.UserControllerInterface,
	controller.TenantControllerInterface,
	controller.ServiceControllerInterface,
	controller.BookingControllerInterface,
	controller.DashboardControllerInterface,
	controller.ChatControllerInterface,
	controller.PaymentControllerInterface,
) {
	tenantRepo := tenantrepo.NewTenantRepository(database)
	userRepo := userrepo.NewUserRepository(database)
	serviceRepo := servicerepo.NewServiceRepository(database)
	bookingRepo := bookingrepo.NewBookingRepository(database)
	chatRepo := chatrepo.NewChatRepository(database)
	paymentRepo := paymentrepo.NewPaymentRepository(database)

	userService := userapp.NewUserDomainService(userRepo)
	tenantService := tenantapp.NewTenantService(tenantRepo)
	tenantOnboarding := tenantapp.NewTenantOnboardingService(tenantRepo, userRepo)
	serviceService := serviceapp.NewServiceService(serviceRepo)
	bookingService := bookingapp.NewBookingService(bookingRepo, tenantRepo, serviceRepo)
	dashboardService := dashboardapp.NewDashboardService(bookingRepo, serviceRepo)

	openAIClient := ai.NewOpenAIClient(os.Getenv("OPENAI_API_KEY"))
	chatService := chatapp.NewChatService(chatRepo, tenantRepo, openAIClient)

	stripeClient := payment.NewStripeClient()
	paymentService := paymentapp.NewPaymentService(paymentRepo, stripeClient)

	return controller.NewUserControllerInterface(userService),
		controller.NewTenantController(tenantService, tenantOnboarding),
		controller.NewServiceController(serviceService),
		controller.NewBookingController(bookingService),
		controller.NewDashboardController(dashboardService),
		controller.NewChatController(chatService),
		controller.NewPaymentController(paymentService)
}
