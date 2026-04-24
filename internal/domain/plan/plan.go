package plan

import (
	"time"
)

type Plan struct {
	ID              string    `json:"id" bson:"_id"`
	Name           string    `json:"name" bson:"name"` // "Free", "Basic", "Pro"
	Description    string    `json:"description" bson:"description"`
	Price          float64   `json:"price" bson:"price"` // em centavos
	Currency       string    `json:"currency" bson:"currency"` // "BRL"
	BillingPeriod  string    `json:"billing_period" bson:"billing_period"` // "monthly", "yearly"
	MaxBookings    int       `json:"max_bookings" bson:"max_bookings"` // -1 = unlimited
	MaxServices   int       `json:"max_services" bson:"max_services"` // -1 = unlimited
	MaxEmployees  int       `json:"max_employees" bson:"max_employees"` // -1 = unlimited
	HasAI          bool      `json:"has_ai" bson:"has_ai"`
	HasWhatsApp   bool      `json:"has_whatsapp" bson:"has_whatsapp"`
	HasCustomDomain bool   `json:"has_custom_domain" bson:"has_custom_domain"`
	HasAnalytics  bool      `json:"has_analytics" bson:"has_analytics"`
	HasPrioritySupport bool `json:"has_priority_support" bson:"has_priority_support"`
	HasWhiteLabel  bool      `json:"has_white_label" bson:"has_white_label"`
	IsActive       bool     `json:"is_active" bson:"is_active"`
	DisplayOrder   int      `json:"display_order" bson:"display_order"`
	CreatedAt      time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" bson:"updated_at"`
}

type Subscription struct {
	ID             string    `json:"id" bson:"_id"`
	TenantID       string    `json:"tenant_id" bson:"tenant_id"`
	PlanID         string    `json:"plan_id" bson:"plan_id"`
	PlanName       string    `json:"plan_name" bson:"plan_name"`
	Status         string    `json:"status" bson:"status"` // active, canceled, past_due, trialing
	CurrentPeriodStart time.Time `json:"current_period_start" bson:"current_period_start"`
	CurrentPeriodEnd   time.Time `json:"current_period_end" bson:"current_period_end"`
	StripeSubscriptionID string `json:"stripe_subscription_id" bson:"stripe_subscription_id"`
	StripeCustomerID   string `json:"stripe_customer_id" bson:"stripe_customer_id"`
	CancelAtPeriodEnd   bool    `json:"cancel_at_period_end" bson:"cancel_at_period_end"`
	TrialStart          *time.Time `json:"trial_start,omitempty" bson:"trial_start,omitempty"`
	TrialEnd            *time.Time `json:"trial_end,omitempty" bson:"trial_end,omitempty"`
	CreatedAt          time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bson:"updated_at"`
}

type Payment struct {
	ID                string    `json:"id" bson:"_id"`
	TenantID          string    `json:"tenant_id" bson:"tenant_id"`
	SubscriptionID   string    `json:"subscription_id" bson:"subscription_id"`
	StripePaymentID  string    `json:"stripe_payment_id" bson:"stripe_payment_id"`
	StripeInvoiceID  string    `json:"stripe_invoice_id" bson:"stripe_invoice_id"`
	Amount            float64   `json:"amount" bson:"amount"`
	Currency          string    `json:"currency" bson:"currency"`
	Status            string    `json:"status" bson:"status"` // succeeded, failed, pending, refunded
	PaymentMethod     string    `json:"payment_method" bson:"payment_method"` // card, pix
	ReceiptURL        string    `json:"receipt_url,omitempty" bson:"receipt_url,omitempty"`
	PaidAt            *time.Time `json:"paid_at,omitempty" bson:"paid_at,omitempty"`
	CreatedAt         time.Time `json:"created_at" bson:"created_at"`
}

func GetDefaultPlans() []Plan {
	return []Plan{
		{
			ID:              "free",
			Name:            "Free",
			Description:     "Perfeito para testar o sistema",
			Price:           0,
			Currency:        "BRL",
			BillingPeriod:   "monthly",
			MaxBookings:     50,
			MaxServices:     1,
			MaxEmployees:   1,
			HasAI:            false,
			HasWhatsApp:     false,
			HasCustomDomain:  false,
			HasAnalytics:    false,
			HasPrioritySupport: false,
			HasWhiteLabel:   false,
			IsActive:        true,
			DisplayOrder:   1,
		},
		{
			ID:              "basic",
			Name:            "Basic",
			Description:     "Ideal para negócios em crescimento",
			Price:           2900,
			Currency:        "BRL",
			BillingPeriod:   "monthly",
			MaxBookings:     200,
			MaxServices:     -1,
			MaxEmployees:    3,
			HasAI:            true,
			HasWhatsApp:     false,
			HasCustomDomain: false,
			HasAnalytics:    true,
			HasPrioritySupport: false,
			HasWhiteLabel:   false,
			IsActive:        true,
			DisplayOrder:   2,
		},
		{
			ID:              "pro",
			Name:            "Pro",
			Description:     "Solução completa para seu negócio",
			Price:           7900,
			Currency:        "BRL",
			BillingPeriod:   "monthly",
			MaxBookings:     -1,
			MaxServices:     -1,
			MaxEmployees:   -1,
			HasAI:            true,
			HasWhatsApp:     true,
			HasCustomDomain: true,
			HasAnalytics:    true,
			HasPrioritySupport: true,
			HasWhiteLabel:   false,
			IsActive:        true,
			DisplayOrder:   3,
		},
		{
			ID:              "enterprise",
			Name:            "Enterprise",
			Description:     "Para grandes equipes e franquias",
			Price:           19900,
			Currency:        "BRL",
			BillingPeriod:   "monthly",
			MaxBookings:     -1,
			MaxServices:     -1,
			MaxEmployees:   -1,
			HasAI:            true,
			HasWhatsApp:     true,
			HasCustomDomain: true,
			HasAnalytics:    true,
			HasPrioritySupport: true,
			HasWhiteLabel:   true,
			IsActive:        true,
			DisplayOrder:   4,
		},
	}
}