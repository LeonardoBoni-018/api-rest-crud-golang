package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	bookingapp "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/booking"
	bookingDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/response"
)

type BookingControllerInterface interface {
	CreateBooking(c *gin.Context)
	ListBookings(c *gin.Context)
	GetBookingByID(c *gin.Context)
	GetAvailability(c *gin.Context)
	GetServicesByTenantSlug(c *gin.Context)
	GetAvailabilityByTenantSlug(c *gin.Context)
	CreateBookingForTenantSlug(c *gin.Context)
	UpdateBookingStatus(c *gin.Context)
	GetTenantPublicInfo(c *gin.Context)
	CancelBookingPublic(c *gin.Context)
}

type bookingController struct {
	booking bookingapp.BookingService
}

func NewBookingController(service bookingapp.BookingService) BookingControllerInterface {
	return &bookingController{booking: service}
}

func (bc *bookingController) CreateBooking(c *gin.Context) {
	var req request.BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	tenantID := c.GetString("tenant_id")
	domainBooking := &bookingDomain.Booking{
		TenantID:      tenantID,
		ServiceID:     req.ServiceID,
		Date:          req.Date,
		TimeSlot:      req.TimeSlot,
		Duration:      req.Duration,
		Status:        "pending",
		CustomerName:  req.CustomerName,
		CustomerPhone: req.CustomerPhone,
		CustomerEmail: req.CustomerEmail,
		Notes:         req.Notes,
	}

	created, err := bc.booking.CreateBooking(domainBooking)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, response.BookingResponse{Booking: created})
}

func (bc *bookingController) ListBookings(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	bookings, err := bc.booking.FindBookingsByTenantID(tenantID)

	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response.BookingsResponse{Bookings: bookings})
}

func (bc *bookingController) GetBookingByID(c *gin.Context) {
	bookingID := c.Param("bookingId")
	bookingModel, err := bc.booking.FindBookingByID(bookingID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response.BookingResponse{Booking: bookingModel})
}

func (bc *bookingController) GetAvailability(c *gin.Context) {
	var req request.AvailabilityRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("invalid date format, expected YYYY-MM-DD"))
		return
	}

	tenantID := c.GetString("tenant_id")
	slots, restErr := bc.booking.GetAvailableSlots(tenantID, req.ServiceID, date)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response.AvailabilityResponse{
		ServiceID: req.ServiceID,
		Date:      req.Date,
		Slots:     slots,
	})
}

func (bc *bookingController) GetServicesByTenantSlug(c *gin.Context) {
	tenantSlug := c.Param("slug")
	services, err := bc.booking.GetServicesByTenantSlug(tenantSlug)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response.ServicesResponse{Services: services})
}

func (bc *bookingController) GetAvailabilityByTenantSlug(c *gin.Context) {
	tenantSlug := c.Param("slug")
	serviceID := c.Param("serviceId")
	dateStr := c.Query("date")
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("invalid date format, expected YYYY-MM-DD"))
		return
	}

	slots, restErr := bc.booking.GetAvailabilityByTenantSlug(tenantSlug, serviceID, date)
	if restErr != nil {
		c.JSON(restErr.Code, restErr)
		return
	}

	c.JSON(http.StatusOK, response.AvailabilityResponse{
		ServiceID: serviceID,
		Date:      dateStr,
		Slots:     slots,
	})
}

func (bc *bookingController) CreateBookingForTenantSlug(c *gin.Context) {
	tenantSlug := c.Param("slug")
	var req request.BookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	booking := &bookingDomain.Booking{
		ServiceID:     req.ServiceID,
		Date:          req.Date,
		TimeSlot:      req.TimeSlot,
		Duration:      req.Duration,
		Status:        "pending",
		CustomerName:  req.CustomerName,
		CustomerPhone: req.CustomerPhone,
		CustomerEmail: req.CustomerEmail,
		Notes:         req.Notes,
	}

	created, err := bc.booking.CreateBookingForTenantSlug(tenantSlug, booking)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusCreated, response.BookingResponse{Booking: created})
}

func (bc *bookingController) UpdateBookingStatus(c *gin.Context) {
	bookingID := c.Param("bookingId")
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	tenantID := c.GetString("tenant_id")
	updatedBooking, err := bc.booking.UpdateBookingStatus(tenantID, bookingID, req.Status)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response.BookingResponse{Booking: updatedBooking})
}

func (bc *bookingController) GetTenantPublicInfo(c *gin.Context) {
	tenantSlug := c.Param("slug")

	tenant, err := bc.booking.GetTenantPublicInfo(tenantSlug)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response.ConvertTenantToPublicInfo(tenant))
}

func (bc *bookingController) CancelBookingPublic(c *gin.Context) {
	bookingID := c.Param("bookingId")
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError("token is required"))
		return
	}

	updated, err := bc.booking.CancelBookingWithToken(bookingID, token)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.JSON(http.StatusOK, response.BookingResponse{Booking: updated})
}
