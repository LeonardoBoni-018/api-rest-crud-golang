package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
	tenantController controller.TenantControllerInterface,
	serviceController controller.ServiceControllerInterface,
	bookingController controller.BookingControllerInterface,
	dashboardController controller.DashboardControllerInterface,
) {

	r.GET("/getUserById/:userId", user.VerifyTokenMiddleware, userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", user.VerifyTokenMiddleware, userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", user.VerifyTokenMiddleware, userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", user.VerifyTokenMiddleware, userController.DeleteUser)
	r.POST("/login", userController.LoginUser)

	r.POST("/tenants/register", tenantController.RegisterTenant)
	r.GET("/tenants/me", user.VerifyTokenMiddleware, tenantController.GetTenantMe)
	r.PATCH("/tenants/settings", user.VerifyTokenMiddleware, tenantController.UpdateTenantSettings)

	r.POST("/services", user.VerifyTokenMiddleware, serviceController.CreateService)
	r.GET("/services", user.VerifyTokenMiddleware, serviceController.ListServices)
	r.PUT("/services/:serviceId", user.VerifyTokenMiddleware, serviceController.UpdateService)
	r.DELETE("/services/:serviceId", user.VerifyTokenMiddleware, serviceController.DeleteService)

	r.POST("/bookings", user.VerifyTokenMiddleware, bookingController.CreateBooking)
	r.GET("/bookings", user.VerifyTokenMiddleware, bookingController.ListBookings)
	r.GET("/bookings/:bookingId", user.VerifyTokenMiddleware, bookingController.GetBookingByID)
	r.GET("/bookings/availability", user.VerifyTokenMiddleware, bookingController.GetAvailability)
	r.PUT("/bookings/:bookingId/status", user.VerifyTokenMiddleware, bookingController.UpdateBookingStatus)

	r.GET("/tenants/:slug/services", bookingController.GetServicesByTenantSlug)
	r.GET("/tenants/:slug/services/:serviceId/availability", bookingController.GetAvailabilityByTenantSlug)
	r.POST("/tenants/:slug/bookings", bookingController.CreateBookingForTenantSlug)

	r.GET("/dashboard/metrics", user.VerifyTokenMiddleware, dashboardController.GetMetrics)
}
