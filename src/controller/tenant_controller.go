package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/internal/application/tenant"
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	userDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/request"
	"github.com/LeonardoBoni-018/api-rest-crud-golang/src/controller/model/response"
)

type TenantControllerInterface interface {
	RegisterTenant(c *gin.Context)
	GetTenantMe(c *gin.Context)
	UpdateTenantSettings(c *gin.Context)
}

type tenantController struct {
	tenantService     tenant.TenantService
	onboardingService tenant.TenantOnboardingService
}

func NewTenantController(
	tenantService tenant.TenantService,
	onboardingService tenant.TenantOnboardingService,
) TenantControllerInterface {
	return &tenantController{
		tenantService:     tenantService,
		onboardingService: onboardingService,
	}
}

func (tc *tenantController) RegisterTenant(c *gin.Context) {
	var req request.TenantRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	tenantModel := &tenantDomain.Tenant{
		Name:  req.Name,
		Slug:  req.Slug,
		Email: req.Email,
		Phone: req.Phone,
		Settings: tenantDomain.Settings{
			BusinessType:    req.BusinessType,
			WorkingDays:     req.WorkingDays,
			WorkingHours:    req.WorkingHours,
			BookingInterval: req.BookingInterval,
			MaxAdvanceDays:  req.MaxAdvanceDays,
		},
	}

	owner := &userDomain.User{
		Email:    req.OwnerEmail,
		Password: req.OwnerPassword,
		Name:     req.OwnerName,
		Phone:    req.OwnerPhone,
	}

	createdTenant, createdOwner, token, err := tc.onboardingService.OnboardTenant(tenantModel, owner)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.Header("Authorization", token)
	c.JSON(http.StatusCreated, response.TenantOnboardResponse{
		Tenant: createdTenant,
		Owner:  createdOwner,
		Token:  token,
	})
}

func (tc *tenantController) GetTenantMe(c *gin.Context) {
	tenantID := c.GetString("tenant_id")
	fmt.Printf("DEBUG - tenant_id from context: '%s'\n", tenantID)
	tenantModel, err := tc.tenantService.FindTenantByID(tenantID)
	if err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.JSON(http.StatusOK, response.TenantResponse{Tenant: tenantModel})
}

func (tc *tenantController) UpdateTenantSettings(c *gin.Context) {
	var req request.TenantSettingsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, rest_err.NewBadRequestError(err.Error()))
		return
	}

	tenantID := c.GetString("tenant_id")
	tenantModel := &tenantDomain.Tenant{
		Settings: tenantDomain.Settings{
			BusinessType:    req.BusinessType,
			WorkingDays:     req.WorkingDays,
			WorkingHours:    req.WorkingHours,
			BookingInterval: req.BookingInterval,
			MaxAdvanceDays:  req.MaxAdvanceDays,
		},
	}

	if err := tc.tenantService.UpdateTenant(tenantID, tenantModel); err != nil {
		c.JSON(err.Code, err)
		return
	}
	c.Status(http.StatusOK)
}
