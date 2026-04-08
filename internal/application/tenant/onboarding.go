package tenant

import (
	"fmt"
	"time"

	"github.com/LeonardoBoni-018/api-rest-crud-golang/configuration/rest_err"
	tenantDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/tenant"
	userDomain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/user"
	tenantRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/tenant/repository"
	userRepo "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/infrastructure/database/user/repository"

)

type TenantOnboardingService interface {
	OnboardTenant(tenant *tenantDomain.Tenant, owner *userDomain.User) (*tenantDomain.Tenant, *userDomain.User, string, *rest_err.RestErr)
}

type tenantOnboardingService struct {
	tenantRepository tenantRepo.TenantRepository
	userRepository   userRepo.UserRepository
}

func NewTenantOnboardingService(
	tenantRepo tenantRepo.TenantRepository,
	userRepo userRepo.UserRepository,
) TenantOnboardingService {
	return &tenantOnboardingService{tenantRepository: tenantRepo, userRepository: userRepo}
}

func (s *tenantOnboardingService) OnboardTenant(
	tenantModel *tenantDomain.Tenant,
	owner *userDomain.User,
) (*tenantDomain.Tenant, *userDomain.User, string, *rest_err.RestErr) {

	tenantModel.Status = "active"
	if tenantModel.Plan == "" {
		tenantModel.Plan = "free"
	}

	tenantModel.CreatedAt = time.Now()
	tenantModel.UpdatedAt = time.Now()

	createdTenant, err := s.tenantRepository.CreateTenant(tenantModel)
	if err != nil {
		return nil, nil, "", err
	}

	fmt.Printf("DEBUG - createdTenant.ID: '%s'\n", createdTenant.ID)

	owner.TenantID = createdTenant.ID
	owner.Role = "owner"
	owner.CreatedAt = time.Now()

	fmt.Printf("DEBUG - owner.TenantID before encrypt: '%s'\n", owner.TenantID)

	// Criptografar a senha antes de salvar
	owner.EncryptPassword()

	// CreateUser agora aceita owner porque *User implementa UserDomainInterface
	savedOwnerInterface, terr := s.userRepository.CreateUser(owner)
	if terr != nil {
		return nil, nil, "", terr
	}

	// Gerar token
	token, terr := savedOwnerInterface.GenerateToken()
	if terr != nil {
		return nil, nil, "", terr
	}

	// Type assertion para converter de interface para *User
	savedOwner, ok := savedOwnerInterface.(*userDomain.User)
	fmt.Printf("DEBUG - savedOwner.TenantID after save: '%s'\n", savedOwner.TenantID)
	if !ok {
		return nil, nil, "", rest_err.NewInternalServerError("error converting user interface")
	}

	return createdTenant, savedOwner, token, nil
}
