package request

type BookingRequest struct {
	ServiceID     string `json:"service_id" binding:"required"`
	Date          string `json:"date" binding:"required"`
	TimeSlot      string `json:"time_slot" binding:"required"`
	Duration      int    `json:"duration" binding:"required,min=1"`
	CustomerName  string `json:"customer_name" binding:"required"`
	CustomerPhone string `json:"customer_phone" binding:"required"`
	CustomerEmail string `json:"customer_email" binding:"omitempty,email"`
	Notes         string `json:"notes"`
}
