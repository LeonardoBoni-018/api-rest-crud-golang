package response

type AvailabilityResponse struct {
	ServiceID string   `json:"service_id"`
	Date      string   `json:"date"`
	Slots     []string `json:"slots"`
}
