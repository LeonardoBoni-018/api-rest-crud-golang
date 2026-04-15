package request

type AvailabilityRequest struct {
	ServiceID string `form:"service_id" binding:"required"`
	Date      string `form:"date" binding:"required"`
}
