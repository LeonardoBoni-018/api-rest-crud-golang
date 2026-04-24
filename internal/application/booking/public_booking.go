package booking

import (
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
)

func (s *bookingService) GetTenantPublicInfo(slug string) (*tenant.Tenant, *rest_err.RestErr) {
	tenant, err := s.tenantRepository.FindTenantBySlug(slug)
	if err != nil {
		return nil, err
	}

	if tenant.Status != "active" {
		return nil, rest_err.NewBadRequestError("business is not currently active")
	}

	return tenant, nil
}

func (s *bookingService) CancelBookingWithToken(bookingID, cancelToken string) (*domain.Booking, *rest_err.RestErr) {
	booking, err := s.bookingRepository.FindBookingByID(bookingID)
	if err != nil {
		return nil, err
	}

	if booking.CancelToken == "" {
		return nil, rest_err.NewBadRequestError("this booking cannot be cancelled online")
	}

	if booking.CancelToken != cancelToken {
		return nil, rest_err.NewBadRequestError("invalid cancellation token")
	}

	if booking.Status == "cancelled" {
		return nil, rest_err.NewBadRequestError("booking is already cancelled")
	}

	if booking.Status == "completed" {
		return nil, rest_err.NewBadRequestError("cannot cancel a completed booking")
	}

	tenant, terr := s.tenantRepository.FindTenantById(booking.TenantID)
	if terr != nil {
		return nil, terr
	}

	serviceItem, serr := s.serviceRepository.FindServiceByID(booking.ServiceID)
	if serr != nil {
		return nil, serr
	}

	booking.Status = "cancelled"
	booking.UpdatedAt = time.Now()

	updated, err := s.bookingRepository.UpdateBooking(booking)
	if err != nil {
		return nil, err
	}

	if updated.CustomerEmail != "" {
		go s.notificationService.SendBookingCancellation(updated, tenant, serviceItem)
	}

	return updated, nil
}