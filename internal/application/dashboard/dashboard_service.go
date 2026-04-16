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
