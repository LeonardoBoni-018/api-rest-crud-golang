package controller

import (
	"net/http"

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
	domain := &bookingDomain.Booking{
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

	created, err := bc.booking.CreateBooking(domain)
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
