package request

type StartChatRequest struct {
	CustomerID string `json:"customer_id"`
	Message    string `json:"message" binding:"required"`
}
