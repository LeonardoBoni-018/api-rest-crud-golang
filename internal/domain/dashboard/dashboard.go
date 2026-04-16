package dashboard

type ServiceMetric struct {
	ServiceID     string  `json:"service_id"`
	Name          string  `json:"name"`
	BookingsCount int     `json:"bookings_count"`
	Revenue       float64 `json:"revenue"`
}

type DashboardMetrics struct {
	TotalBookings     int             `json:"total_bookings"`
	ConfirmedBookings int             `json:"confirmed_bookings"`
	CancelledBookings int             `json:"cancelled_bookings"`
	Revenue           float64         `json:"revenue"`
	UpcomingBookings  int             `json:"upcoming_bookings"`
	TopServices       []ServiceMetric `json:"top_services"`
}
