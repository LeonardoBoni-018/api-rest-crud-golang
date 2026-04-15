package booking

import (
	"fmt"
	"strings"
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/booking"
)

const (
	defaultWorkingHours    = "09:00-18:00"
	defaultBookingInterval = 30
	dateLayout             = "2006-01-02"
	timeLayout             = "15:04"
)

func (s *bookingService) GetAvailableSlots(tenantID, serviceID string, date time.Time) ([]string, *rest_err.RestErr) {
	tenant, err := s.tenantRepository.FindTenantById(tenantID)
	if err != nil {
		return nil, err
	}

	if !isWorkingDayAllowed(tenant.Settings.WorkingDays, date) {
		return []string{}, nil
	}

	if tenant.Settings.MaxAdvanceDays > 0 {
		maxDate := time.Now().AddDate(0, 0, tenant.Settings.MaxAdvanceDays)
		if date.After(maxDate) {
			return nil, rest_err.NewBadRequestError(
				fmt.Sprintf("booking date must be within %d days", tenant.Settings.MaxAdvanceDays),
			)
		}
	}

	service, err := s.serviceRepository.FindServiceByID(serviceID)
	if err != nil {
		return nil, err
	}

	if service.TenantID != tenantID {
		return nil, rest_err.NewBadRequestError("service does not belong to tenant")
	}

	interval := tenant.Settings.BookingInterval
	if interval == 0 {
		interval = defaultBookingInterval
	}

	duration := service.Duration
	if duration <= 0 {
		duration = interval
	}

	workingHours := strings.TrimSpace(tenant.Settings.WorkingHours)
	if workingHours == "" {
		workingHours = defaultWorkingHours
	}

	startTime, endTime, err := parseWorkingHours(workingHours, date)
	if err != nil {
		return nil, rest_err.NewBadRequestError("invalid working hours format")
	}

	bookings, err := s.bookingRepository.FindBookingsByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	sameDaysBookings := filterBookingsByDate(bookings, date.Format(dateLayout))
	slots := buildSlots(startTime, endTime, duration, interval)

	return filterBookedSlots(slots, sameDaysBookings, duration), nil
}

func parseWorkingHours(workingHours string, date time.Time) (time.Time, time.Time, *rest_err.RestErr) {
	parts := strings.Split(workingHours, "-")
	if len(parts) != 2 {
		return time.Time{}, time.Time{}, rest_err.NewBadRequestError("invalid working hours format")
	}

	start, err := time.Parse(timeLayout, strings.TrimSpace(parts[0]))
	if err != nil {
		return time.Time{}, time.Time{}, rest_err.NewBadRequestError("invalid start time format")
	}

	end, err := time.Parse(timeLayout, strings.TrimSpace(parts[1]))
	if err != nil {
		return time.Time{}, time.Time{}, rest_err.NewBadRequestError("invalid end time format")
	}

	startTime := time.Date(date.Year(), date.Month(), date.Day(), start.Hour(), start.Minute(), 0, 0, date.Location())
	endTime := time.Date(date.Year(), date.Month(), date.Day(), end.Hour(), end.Minute(), 0, 0, date.Location())

	if !endTime.After(startTime) {
		return time.Time{}, time.Time{}, rest_err.NewBadRequestError("end time must be after start time")
	}

	return startTime, endTime, nil
}

func isWorkingDayAllowed(days []string, date time.Time) bool {
	if len(days) == 0 {
		return true
	}

	weekDay := strings.ToLower(date.Weekday().String())
	shortWeekDay := weekDay[:3]

	for _, day := range days {
		normalizedDay := strings.ToLower(strings.TrimSpace(day))
		if normalizedDay == weekDay || normalizedDay == shortWeekDay {
			return true
		}
	}
	return false
}

func buildSlots(start, end time.Time, durationMinutes, intervalMinutes int) []time.Time {
	var slots []time.Time
	for current := start; ; current = current.Add(time.Duration(intervalMinutes) * time.Minute) {
		if current.Add(time.Duration(durationMinutes) * time.Minute).After(end) {
			break
		}
		slots = append(slots, current)
	}
	return slots
}

func filterBookingsByDate(bookings []*domain.Booking, date string) []*domain.Booking {
	var filtered []*domain.Booking
	for _, booking := range bookings {
		if booking.Date == date && !strings.EqualFold(booking.Status, "cancelled") {
			filtered = append(filtered, booking)
		}
	}
	return filtered
}

func filterBookedSlots(slots []time.Time, bookings []*domain.Booking, durationMinutes int) []string {
	var available []string
	for _, slot := range slots {
		slotEnd := slot.Add(time.Duration(durationMinutes) * time.Minute)

		if overlapsAny(slot, slotEnd, bookings) {
			continue
		}

		available = append(available, slot.Format(timeLayout))
	}
	return available
}

func overlapsAny(slotStart, slotEnd time.Time, bookings []*domain.Booking) bool {
	for _, booking := range bookings {
		bookingStart, err := time.ParseInLocation(timeLayout, booking.TimeSlot, slotStart.Location())
		if err != nil {
			continue
		}

		bookedStart := time.Date(slotStart.Year(), slotStart.Month(), slotStart.Day(), bookingStart.Hour(), bookingStart.Minute(), 0, 0, slotStart.Location())
		bookedEnd := bookedStart.Add(time.Duration(booking.Duration) * time.Minute)
		if booking.Duration <= 0 {
			bookedEnd = bookedStart.Add(time.Duration(defaultBookingInterval) * time.Minute)
		}

		if slotStart.Before(bookedEnd) && bookedStart.Before(slotEnd) {
			return true
		}
	}
	return false
}
