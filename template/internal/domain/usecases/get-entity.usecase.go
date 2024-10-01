package usecases

import (
	"template/internal/domain/entities"
	"template/internal/domain/gateways"
)

// GetEntityUseCase defines the use case for retrieving an entity
type GetEntityUseCase struct {
	gateway gateways.EntityGateway
}

// NewGetEntityUseCase creates a new get entity use case
func NewGetEntityUseCase(gateway gateways.EntityGateway) *GetEntityUseCase {
	return &GetEntityUseCase{gateway: gateway}
}

// Execute retrieves an entity by ID
func (uc *GetEntityUseCase) Execute(id string) (entities.Entity, error) {
	return uc.gateway.Get(id)
}
