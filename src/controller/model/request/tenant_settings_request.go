package request

type TenantSettingsRequest struct {
	BusinessType    string   `json:"business_type" binding:"required"`
	WorkingDays     []string `json:"working_days" binding:"required"`
	WorkingHours    string   `json:"working_hours" binding:"required"`
	BookingInterval int      `json:"booking_interval" binding:"required"`
	MaxAdvanceDays  int      `json:"max_advance_days" binding:"required"`
}
