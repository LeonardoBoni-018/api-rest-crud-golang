package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SubscriptionEntity struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TenantID       primitive.ObjectID `json:"tenant_id" bson:"tenant_id"`
	PlanID         string             `json:"plan_id" bson:"plan_id"`
	PlanName       string             `json:"plan_name" bson:"plan_name"`
	Status         string             `json:"status" bson:"status"`
	CurrentPeriodStart time.Time       `json:"current_period_start" bson:"current_period_start"`
	CurrentPeriodEnd   time.Time       `json:"current_period_end" bson:"current_period_end"`
	StripeSubscriptionID string       `json:"stripe_subscription_id" bson:"stripe_subscription_id"`
	StripeCustomerID   string         `json:"stripe_customer_id" bson:"stripe_customer_id"`
	CancelAtPeriodEnd   bool           `json:"cancel_at_period_end" bson:"cancel_at_period_end"`
	TrialStart          *time.Time     `json:"trial_start,omitempty" bson:"trial_start,omitempty"`
	TrialEnd            *time.Time     `json:"trial_end,omitempty" bson:"trial_end,omitempty"`
	CreatedAt          time.Time     `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time     `json:"updated_at" bson:"updated_at"`
}

type PaymentEntity struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	TenantID          primitive.ObjectID `json:"tenant_id" bson:"tenant_id"`
	SubscriptionID  primitive.ObjectID `json:"subscription_id" bson:"subscription_id"`
	StripePaymentID  string          `json:"stripe_payment_id" bson:"stripe_payment_id"`
	StripeInvoiceID  string          `json:"stripe_invoice_id" bson:"stripe_invoice_id"`
	Amount            float64         `json:"amount" bson:"amount"`
	Currency         string          `json:"currency" bson:"currency"`
	Status           string          `json:"status" bson:"status"`
	PaymentMethod    string          `json:"payment_method" bson:"payment_method"`
	ReceiptURL       string          `json:"receipt_url,omitempty" bson:"receipt_url,omitempty"`
	PaidAt           *time.Time      `json:"paid_at,omitempty" bson:"paid_at,omitempty"`
	CreatedAt        time.Time       `json:"created_at" bson:"created_at"`
}