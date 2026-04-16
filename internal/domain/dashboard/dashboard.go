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

type DashboardCalendarEvent struct {
	BookingID     string  `json:"booking_id"`
	Date          string  `json:"date"`
	TimeSlot      string  `json:"time_slot"`
	Duration      int     `json:"duration"`
	Status        string  `json:"status"`
	ServiceName   string  `json:"service_name"`
	CustomerName  string  `json:"customer_name"`
	CustomerPhone string  `json:"customer_phone"`
	Revenue       float64 `json:"revenue"`
}

type DashboardCalendar struct {
	Events []DashboardCalendarEvent `json:"events"`
}

type ReportStatus struct {
	Status string `json:"status"`
	Count  int    `json:"count"`
}

type ReportRevenueByService struct {
	ServiceID     string  `json:"service_id"`
	Name          string  `json:"name"`
	BookingsCount int     `json:"bookings_count"`
	Revenue       float64 `json:"revenue"`
}

type DailyBooking struct {
	Date     string  `json:"date"`
	Bookings int     `json:"bookings"`
	Revenue  float64 `json:"revenue"`
}

type DashboardReports struct {
	TotalBookings    int                      `json:"total_bookings"`
	StatusSummary    []ReportStatus           `json:"status_summary"`
	RevenueByService []ReportRevenueByService `json:"revenue_by_service"`
	DailyBookings    []DailyBooking           `json:"daily_bookings"`
}
