package request

type ServiceRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Duration    int     `json:"duration" binding:"required,min=1"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	IsActive    bool    `json:"is_active"`
}
