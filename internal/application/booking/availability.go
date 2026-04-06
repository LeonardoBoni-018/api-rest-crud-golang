package booking

import (
	"time"
)

type AvailabilityService struct {
	// TODO: add repository dependencies
}

func NewAvailabilityService() *AvailabilityService {
	return &AvailabilityService{}
}

func (s *AvailabilityService) GetAvailableSlots(tenantID string, date time.Time) ([]string, error) {
	// TODO: implementar lógica de disponibilidade:
	// 1. carregar configurações do tenant
	// 2. gerar slots com base em working_hours e booking_interval
	// 3. buscar agendamentos existentes
	// 4. filtrar os slots ocupados
	return []string{}, nil
}
