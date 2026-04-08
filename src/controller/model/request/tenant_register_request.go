package request

type TenantRegisterRequest struct {
	Name            string   `json:"name" binding:"required"`
	Slug            string   `json:"slug" binding:"required"`
	Email           string   `json:"email" binding:"required,email"`
	Phone           string   `json:"phone" binding:"required"`
	BusinessType    string   `json:"business_type" binding:"required"`
	WorkingDays     []string `json:"working_days" binding:"required"`
	WorkingHours    string   `json:"working_hours" binding:"required"`
	BookingInterval int      `json:"booking_interval" binding:"required"`
	MaxAdvanceDays  int      `json:"max_advance_days" binding:"required"`

	OwnerName     string `json:"owner_name" binding:"required"`
	OwnerEmail    string `json:"owner_email" binding:"required,email"`
	OwnerPassword string `json:"owner_password" binding:"required"`
	OwnerPhone    string `json:"owner_phone" binding:"required"`
}
