package dashboard

import (
	"sort"
	"strings"
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	dashboarddomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/dashboard"
	servicedomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"
)

const (
	dateLayout = "2006-01-02"
)

func (s *dashboardService) GetMetrics(tenantID string) (*dashboarddomain.DashboardMetrics, *rest_err.RestErr) {
	bookings, err := s.bookingRepository.FindBookingsByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	services, err := s.serviceRepository.FindServicesByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	serviceMap := make(map[string]*servicedomain.Service, len(services))
	for _, service := range services {
		serviceMap[service.ID] = service
	}

	metrics := &dashboarddomain.DashboardMetrics{}
	today := time.Now().Format(dateLayout)
	metrics.TopServices = []dashboarddomain.ServiceMetric{}
	serviceStats := make(map[string]*dashboarddomain.ServiceMetric)

	for _, booking := range bookings {
		metrics.TotalBookings++

		status := strings.ToLower(strings.TrimSpace(booking.Status))
		if status == "confirmed" || status == "completed" {
			metrics.ConfirmedBookings++
		}
		if status == "cancelled" {
			metrics.CancelledBookings++
		}

		if booking.Date >= today && status != "cancelled" {
			metrics.UpcomingBookings++
		}

		servicePrice := getServicePrice(serviceMap, booking.ServiceID)
		if status == "confirmed" || status == "completed" {
			metrics.Revenue += servicePrice
		}

		if _, ok := serviceStats[booking.ServiceID]; !ok {
			serviceStats[booking.ServiceID] = &dashboarddomain.ServiceMetric{
				ServiceID: booking.ServiceID,
				Name:      getServiceName(serviceMap, booking.ServiceID),
			}
		}

		if status != "cancelled" {
			serviceStats[booking.ServiceID].BookingsCount++
			serviceStats[booking.ServiceID].Revenue += servicePrice
		}
	}

	for _, stat := range serviceStats {
		metrics.TopServices = append(metrics.TopServices, *stat)
	}

	sort.Slice(metrics.TopServices, func(i, j int) bool {
		return metrics.TopServices[i].BookingsCount > metrics.TopServices[j].BookingsCount
	})

	if len(metrics.TopServices) > 5 {
		metrics.TopServices = metrics.TopServices[:5]
	}

	return metrics, nil
}

func getServicePrice(serviceMap map[string]*servicedomain.Service, serviceID string) float64 {
	if service, ok := serviceMap[serviceID]; ok {
		return service.Price
	}
	return 0
}

func getServiceName(serviceMap map[string]*servicedomain.Service, serviceID string) string {
	if service, ok := serviceMap[serviceID]; ok {
		return service.Name
	}
	return ""
}

func (s *dashboardService) GetCalendar(tenantID string) (*dashboarddomain.DashboardCalendar, *rest_err.RestErr) {
	bookings, err := s.bookingRepository.FindBookingsByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	services, err := s.serviceRepository.FindServicesByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	serviceMap := make(map[string]*servicedomain.Service, len(services))
	for _, service := range services {
		serviceMap[service.ID] = service
	}

	var events []dashboarddomain.DashboardCalendarEvent
	now := time.Now()

	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	for _, booking := range bookings {
		bookingTime, errParse := time.Parse("2006-01-02 15:04", booking.Date+" "+booking.TimeSlot)
		if errParse != nil {
			continue
		}

		if bookingTime.Before(today) {
			continue
		}

		events = append(events, dashboarddomain.DashboardCalendarEvent{
			BookingID:     booking.ID,
			Date:          booking.Date,
			TimeSlot:      booking.TimeSlot,
			Duration:      booking.Duration,
			Status:        booking.Status,
			ServiceName:   getServiceName(serviceMap, booking.ServiceID),
			CustomerName:  booking.CustomerName,
			CustomerPhone: booking.CustomerPhone,
			Revenue:       getServicePrice(serviceMap, booking.ServiceID),
		})
	}

	sort.Slice(events, func(i, j int) bool {
		left := events[i].Date + " " + events[i].TimeSlot
		right := events[j].Date + " " + events[j].TimeSlot
		return left < right
	})

	return &dashboarddomain.DashboardCalendar{Events: events}, nil
}

func (s *dashboardService) GetReports(tenantID string) (*dashboarddomain.DashboardReports, *rest_err.RestErr) {
	bookings, err := s.bookingRepository.FindBookingsByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	services, err := s.serviceRepository.FindServicesByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	serviceMap := make(map[string]*servicedomain.Service, len(services))
	for _, service := range services {
		serviceMap[service.ID] = service
	}

	statusCounts := map[string]int{}
	revenueByService := map[string]*dashboarddomain.ReportRevenueByService{}
	dailyStats := map[string]*dashboarddomain.DailyBooking{}
	sevenDays := time.Now().AddDate(0, 0, -6)

	for _, booking := range bookings {
		status := strings.ToLower(strings.TrimSpace(booking.Status))
		statusCounts[status]++

		price := getServicePrice(serviceMap, booking.ServiceID)
		if _, ok := revenueByService[booking.ServiceID]; !ok {
			revenueByService[booking.ServiceID] = &dashboarddomain.ReportRevenueByService{
				ServiceID: booking.ServiceID,
				Name:      getServiceName(serviceMap, booking.ServiceID),
			}
		}

		if status != "cancelled" {
			revenueByService[booking.ServiceID].BookingsCount++
			revenueByService[booking.ServiceID].Revenue += price
		}

		bookingDate, errParse := time.Parse(dateLayout, booking.Date)
		if errParse != nil {
			continue
		}

		if bookingDate.After(sevenDays) || bookingDate.Equal(sevenDays) {
			key := bookingDate.Format(dateLayout)
			if _, ok := dailyStats[key]; !ok {
				dailyStats[key] = &dashboarddomain.DailyBooking{Date: key}
			}
			dailyStats[key].Bookings++
			if status != "cancelled" {
				dailyStats[key].Revenue += price
			}
		}
	}

	var statusSummary []dashboarddomain.ReportStatus
	for status, count := range statusCounts {
		statusSummary = append(statusSummary, dashboarddomain.ReportStatus{Status: status, Count: count})
	}

	var revenueSummary []dashboarddomain.ReportRevenueByService
	for _, item := range revenueByService {
		revenueSummary = append(revenueSummary, *item)
	}

	var dailyBookings []dashboarddomain.DailyBooking
	for _, item := range dailyStats {
		dailyBookings = append(dailyBookings, *item)
	}

	sort.Slice(statusSummary, func(i, j int) bool {
		return statusSummary[i].Count > statusSummary[j].Count
	})
	sort.Slice(revenueSummary, func(i, j int) bool {
		return revenueSummary[i].Revenue > revenueSummary[j].Revenue
	})
	sort.Slice(dailyBookings, func(i, j int) bool {
		return dailyBookings[i].Date < dailyBookings[j].Date
	})

	return &dashboarddomain.DashboardReports{
		TotalBookings:    len(bookings),
		StatusSummary:    statusSummary,
		RevenueByService: revenueSummary,
		DailyBookings:    dailyBookings,
	}, nil
}
