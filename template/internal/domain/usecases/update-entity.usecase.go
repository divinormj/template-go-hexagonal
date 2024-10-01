package usecases

import (
	"template/internal/domain/entities"
	"template/internal/domain/gateways"
)

// UpdateEntityUseCase defines the use case for updating an entity
type UpdateEntityUseCase struct {
	gateway gateways.EntityGateway
}

// NewUpdateEntityUseCase creates a new update entity use case
func NewUpdateEntityUseCase(gateway gateways.EntityGateway) *UpdateEntityUseCase {
	return &UpdateEntityUseCase{gateway: gateway}
}

// Execute updates an existing entity
func (uc *UpdateEntityUseCase) Execute(entity entities.Entity) (entities.Entity, error) {
	return uc.gateway.Update(entity)
}
