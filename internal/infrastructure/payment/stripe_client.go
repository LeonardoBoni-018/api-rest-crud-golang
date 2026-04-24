package payment

import (
	"os"
)

type StripeClient struct {
	apiKey        string
	webhookSecret string
	isConfigured  bool
}

func NewStripeClient() *StripeClient {
	apiKey := os.Getenv("STRIPE_API_KEY")
	webhookSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	isConfigured := apiKey != ""

	return &StripeClient{
		apiKey:        apiKey,
		webhookSecret: webhookSecret,
		isConfigured: isConfigured,
	}
}

func (c *StripeClient) IsConfigured() bool {
	return c.isConfigured
}

func (c *StripeClient) CreateCustomer(email, name, tenantID string) (interface{}, error) {
	if !c.isConfigured {
		return nil, nil
	}
	return nil, nil
}

func (c *StripeClient) CreateSubscription(customerID, priceID string) (interface{}, error) {
	if !c.isConfigured {
		return nil, nil
	}
	return nil, nil
}

func (c *StripeClient) CreateCheckoutSession(params *CheckoutParams) (interface{}, error) {
	if !c.isConfigured {
		return nil, nil
	}
	return nil, nil
}

func (c *StripeClient) CancelSubscription(subID string, cancelAtPeriodEnd bool) (interface{}, error) {
	if !c.isConfigured {
		return nil, nil
	}
	return nil, nil
}

func (c *StripeClient) GetSubscription(subID string) (interface{}, error) {
	if !c.isConfigured {
		return nil, nil
	}
	return nil, nil
}

func (c *StripeClient) ConstructWebhookEvent(payload []byte, signature string) (interface{}, error) {
	if !c.isConfigured {
		return nil, nil
	}
	return nil, nil
}

type CheckoutParams struct {
	TenantID        string
	PlanID          string
	PriceID         string
	CustomerID      string
	CustomerEmail   string
	IsSubscription bool
	SuccessURL      string
	CancelURL       string
}

type PlanConfig struct {
	ID            string
	Name          string
	Price         int64
	BillingPeriod string
}

func GetPlanPriceID(planID string) string {
	priceIDs := map[string]string{
		"free":      "",
		"basic":     os.Getenv("STRIPE_PRICE_BASIC_ID"),
		"pro":       os.Getenv("STRIPE_PRICE_PRO_ID"),
		"enterprise": os.Getenv("STRIPE_PRICE_ENTERPRISE_ID"),
	}

	return priceIDs[planID]
}