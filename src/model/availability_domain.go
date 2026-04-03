package model

type AvailabilityDomain struct {
	ID         string `bson:"_id,omitempty" json:"id"`
	BusinessID string `bson:"business_id" json:"business_id"`
	Weekday    int    `bson:"weekday" json:"weekday"`       // 0=domingo..6=sábado
	StartTime  string `bson:"start_time" json:"start_time"` // "09:00"
	EndTime    string `bson:"end_time" json:"end_time"`     // "18:00"
	SlotMin    int    `bson:"slot_min" json:"slot_min"`     // 15/30/60
}
