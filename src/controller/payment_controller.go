package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/payment"
)

type PaymentControllerInterface interface {
	GetPlans(*gin.Context)
	CreateCheckout(*gin.Context)
	GetSubscription(*gin.Context)
	CancelSubscription(*gin.Context)
	HandleWebhook(*gin.Context)
}

type paymentController struct {
	paymentService payment.PaymentService
}

func NewPaymentController(service payment.PaymentService) PaymentControllerInterface {
	return &paymentController{paymentService: service}
}

func (c *paymentController) GetPlans(ctx *gin.Context) {
	plans := c.paymentService.GetPlans()
	if plans == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get plans"})
		return
	}
	ctx.JSON(http.StatusOK, plans)
}

func (c *paymentController) CreateCheckout(ctx *gin.Context) {
	var req payment.CheckoutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		err := rest_err.NewBadRequestError("invalid request body")
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	tenantID := ctx.GetString("tenant_id")
	if tenantID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	req.TenantID = tenantID

	resp, restErr := c.paymentService.CreateCheckoutSession(&req)
	if restErr != nil {
		ctx.JSON(restErr.Code, gin.H{"error": restErr.Message})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (c *paymentController) GetSubscription(ctx *gin.Context) {
	tenantID := ctx.GetString("tenant_id")
	if tenantID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	sub, restErr := c.paymentService.GetSubscriptionByTenantID(tenantID)
	if restErr != nil {
		ctx.JSON(restErr.Code, gin.H{"error": restErr.Message})
		return
	}

	ctx.JSON(http.StatusOK, sub)
}

func (c *paymentController) CancelSubscription(ctx *gin.Context) {
	tenantID := ctx.GetString("tenant_id")
	if tenantID == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	var req struct {
		CancelAtPeriodEnd bool `json:"cancel_at_period_end"`
	}
	ctx.ShouldBindJSON(&req)

	sub, restErr := c.paymentService.CancelSubscription(tenantID, req.CancelAtPeriodEnd)
	if restErr != nil {
		ctx.JSON(restErr.Code, gin.H{"error": restErr.Message})
		return
	}

	ctx.JSON(http.StatusOK, sub)
}

func (c *paymentController) HandleWebhook(ctx *gin.Context) {
	payload, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	signature := ctx.GetHeader("Stripe-Signature")
	restErr := c.paymentService.HandleWebhook(payload, signature)
	if restErr != nil {
		ctx.JSON(restErr.Code, gin.H{"error": restErr.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "received"})
}