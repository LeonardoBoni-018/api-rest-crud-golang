package response

import domain "github.com/LeonardoBoni-018/api-rest-crud-golang/internal/domain/service"

type ServiceResponse struct {
	Service *domain.Service `json:"service"`
}

type ServicesResponse struct {
	Services []*domain.Service `json:"services"`
}
