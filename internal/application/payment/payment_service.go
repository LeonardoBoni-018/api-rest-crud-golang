package payment

import (
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/plan"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/payment/repository"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/payment"
)

type PaymentService interface {
	GetPlans() []plan.Plan
	GetPlanByID(string) (*plan.Plan, *rest_err.RestErr)
	CreateCheckoutSession(*CheckoutRequest) (*CheckoutResponse, *rest_err.RestErr)
	GetSubscriptionByTenantID(string) (*plan.Subscription, *rest_err.RestErr)
	CancelSubscription(string, bool) (*plan.Subscription, *rest_err.RestErr)
	HandleWebhook([]byte, string) *rest_err.RestErr
}

type paymentService struct {
	paymentRepo  repository.PaymentRepository
	stripeClient *payment.StripeClient
}

func NewPaymentService(repo repository.PaymentRepository, stripe *payment.StripeClient) PaymentService {
	return &paymentService{
		paymentRepo:  repo,
		stripeClient: stripe,
	}
}

type CheckoutRequest struct {
	TenantID      string `json:"tenant_id"`
	PlanID        string `json:"plan_id"`
	CustomerEmail string `json:"customer_email"`
	TenantName    string `json:"tenant_name"`
	SuccessURL    string `json:"success_url"`
	CancelURL     string `json:"cancel_url"`
}

type CheckoutResponse struct {
	SessionID    string `json:"session_id"`
	SessionURL  string `json:"session_url"`
	Gateway     string `json:"gateway"`
}

func (s *paymentService) GetPlans() []plan.Plan {
	return plan.GetDefaultPlans()
}

func (s *paymentService) GetPlanByID(planID string) (*plan.Plan, *rest_err.RestErr) {
	plans := plan.GetDefaultPlans()
	for _, p := range plans {
		if p.ID == planID {
			return &p, nil
		}
	}
	return nil, rest_err.NewNotFoundError("plan not found")
}

func (s *paymentService) CreateCheckoutSession(req *CheckoutRequest) (*CheckoutResponse, *rest_err.RestErr) {
	planInfo, err := s.GetPlanByID(req.PlanID)
	if err != nil {
		return nil, err
	}

	if planInfo.Price == 0 {
		sub := &plan.Subscription{
			TenantID:  req.TenantID,
			PlanID:    req.PlanID,
			PlanName:  planInfo.Name,
			Status:    "active",
			CurrentPeriodStart: time.Now(),
			CurrentPeriodEnd:   time.Now().AddDate(0, 1, 0),
		}
		_, err := s.paymentRepo.CreateSubscription(sub)
		if err != nil {
			return nil, err
		}
		return &CheckoutResponse{
			SessionID:   "free",
			SessionURL: "",
			Gateway:    "free",
		}, nil
	}

	var customerID string
	if s.stripeClient.IsConfigured() {
		stripeCustomer, err := s.stripeClient.CreateCustomer(req.CustomerEmail, req.TenantName, req.TenantID)
		if err != nil {
			return nil, rest_err.NewInternalServerError("error creating customer")
		}
		if stripeCustomer != nil {
			if c, ok := stripeCustomer.(interface{ GetID() string }); ok {
				customerID = c.GetID()
			}
		}

		priceID := payment.GetPlanPriceID(req.PlanID)
		if priceID == "" {
			return nil, rest_err.NewBadRequestError("plan not configured for payment")
		}

		stripeSession, err := s.stripeClient.CreateCheckoutSession(&payment.CheckoutParams{
			TenantID:        req.TenantID,
			PlanID:          req.PlanID,
			PriceID:         priceID,
			CustomerID:      customerID,
			CustomerEmail:   req.CustomerEmail,
			IsSubscription:  true,
			SuccessURL:      req.SuccessURL,
			CancelURL:       req.CancelURL,
		})
		if err != nil {
			return nil, rest_err.NewInternalServerError("error creating checkout session")
		}

		if stripeSession == nil {
			return &CheckoutResponse{
				SessionID:   "test",
				SessionURL: "https://checkout.stripe.com/test-mode",
				Gateway:    "stripe",
			}, nil
		}

		if session, ok := stripeSession.(interface{ GetID() string; GetURL() string }); ok {
			return &CheckoutResponse{
				SessionID:   session.GetID(),
				SessionURL: session.GetURL(),
				Gateway:    "stripe",
			}, nil
		}

		return &CheckoutResponse{
			SessionID:   "demo",
			SessionURL: "",
			Gateway:    "stripe",
		}, nil
	}

	return &CheckoutResponse{
		SessionID:   "demo",
		SessionURL: "",
		Gateway:    "demo",
	}, nil
}

func (s *paymentService) GetSubscriptionByTenantID(tenantID string) (*plan.Subscription, *rest_err.RestErr) {
	sub, err := s.paymentRepo.FindSubscriptionByTenantId(tenantID)
	if err != nil {
		return nil, err
	}

	if sub == nil {
		return &plan.Subscription{
			TenantID: tenantID,
			PlanID:  "free",
			PlanName: "Free",
			Status:  "active",
		}, nil
	}

	return sub, nil
}

func (s *paymentService) CancelSubscription(tenantID string, cancelAtPeriodEnd bool) (*plan.Subscription, *rest_err.RestErr) {
	sub, err := s.GetSubscriptionByTenantID(tenantID)
	if err != nil {
		return nil, err
	}

	if sub.StripeSubscriptionID != "" && s.stripeClient.IsConfigured() {
		_, err := s.stripeClient.CancelSubscription(sub.StripeSubscriptionID, cancelAtPeriodEnd)
		if err != nil {
			return nil, rest_err.NewInternalServerError("error canceling subscription")
		}
	}

	sub.CancelAtPeriodEnd = cancelAtPeriodEnd
	if cancelAtPeriodEnd {
		sub.Status = "canceled"
	} else {
		sub.Status = "canceled"
		sub.CurrentPeriodEnd = time.Now()
	}

	_, err = s.paymentRepo.UpdateSubscription(sub.ID, sub)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *paymentService) HandleWebhook(payload []byte, signature string) *rest_err.RestErr {
	return nil
}

func (s *paymentService) handleSubscriptionEvent(event interface{}) {
}

func (s *paymentService) handleInvoiceEvent(event interface{}) {
}

func (s *paymentService) handleCheckoutSessionEvent(event interface{}) {
}