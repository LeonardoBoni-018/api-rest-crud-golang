package converter

import (
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/plan"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/payment/repository/entity"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToEntitySubscription(s *plan.Subscription) *entity.SubscriptionEntity {
	if s == nil {
		return nil
	}
	tenantOID, _ := primitive.ObjectIDFromHex(s.TenantID)
	return &entity.SubscriptionEntity{
		ID:             primitive.ObjectID{},
		TenantID:       tenantOID,
		PlanID:         s.PlanID,
		PlanName:       s.PlanName,
		Status:         s.Status,
		CurrentPeriodStart: s.CurrentPeriodStart,
		CurrentPeriodEnd:   s.CurrentPeriodEnd,
		StripeSubscriptionID: s.StripeSubscriptionID,
		StripeCustomerID:   s.StripeCustomerID,
		CancelAtPeriodEnd:   s.CancelAtPeriodEnd,
		TrialStart:          s.TrialStart,
		TrialEnd:            s.TrialEnd,
		CreatedAt:          s.CreatedAt,
		UpdatedAt:          s.UpdatedAt,
	}
}

func ToDomainSubscription(e *entity.SubscriptionEntity) *plan.Subscription {
	if e == nil {
		return nil
	}
	return &plan.Subscription{
		ID:             e.ID.Hex(),
		TenantID:       e.TenantID.Hex(),
		PlanID:         e.PlanID,
		PlanName:       e.PlanName,
		Status:         e.Status,
		CurrentPeriodStart: e.CurrentPeriodStart,
		CurrentPeriodEnd:   e.CurrentPeriodEnd,
		StripeSubscriptionID: e.StripeSubscriptionID,
		StripeCustomerID:   e.StripeCustomerID,
		CancelAtPeriodEnd:   e.CancelAtPeriodEnd,
		TrialStart:          e.TrialStart,
		TrialEnd:            e.TrialEnd,
		CreatedAt:          e.CreatedAt,
		UpdatedAt:          e.UpdatedAt,
	}
}

func ToEntityPayment(p *plan.Payment) *entity.PaymentEntity {
	if p == nil {
		return nil
	}
	tenantOID, _ := primitive.ObjectIDFromHex(p.TenantID)
	subscriptionOID, _ := primitive.ObjectIDFromHex(p.SubscriptionID)
	var paidAt *time.Time
	if p.PaidAt != nil {
		t := *p.PaidAt
		paidAt = &t
	}
	return &entity.PaymentEntity{
		ID:               primitive.ObjectID{},
		TenantID:         tenantOID,
		SubscriptionID:  subscriptionOID,
		StripePaymentID: p.StripePaymentID,
		StripeInvoiceID: p.StripeInvoiceID,
		Amount:          p.Amount,
		Currency:       p.Currency,
		Status:          p.Status,
		PaymentMethod:  p.PaymentMethod,
		ReceiptURL:     p.ReceiptURL,
		PaidAt:        paidAt,
		CreatedAt:      p.CreatedAt,
	}
}

func ToDomainPayment(e *entity.PaymentEntity) *plan.Payment {
	if e == nil {
		return nil
	}
	var paidAt *time.Time
	if e.PaidAt != nil {
		t := *e.PaidAt
		paidAt = &t
	}
	return &plan.Payment{
		ID:               e.ID.Hex(),
		TenantID:         e.TenantID.Hex(),
		SubscriptionID:  e.SubscriptionID.Hex(),
		StripePaymentID: e.StripePaymentID,
		StripeInvoiceID: e.StripeInvoiceID,
		Amount:         e.Amount,
		Currency:       e.Currency,
		Status:         e.Status,
		PaymentMethod: e.PaymentMethod,
		ReceiptURL:    e.ReceiptURL,
		PaidAt:        paidAt,
		CreatedAt:      e.CreatedAt,
	}
}