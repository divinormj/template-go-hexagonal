// Package usecases contains the application use cases.
package usecases

import (
	"template/internal/domain/entities"
	"template/internal/domain/gateways"
)

// CreateEntityUseCase defines the use case for creating an entity
type CreateEntityUseCase struct {
	gateway gateways.EntityGateway
}

// NewCreateEntityUseCase creates a new create entity use case
func NewCreateEntityUseCase(gateway gateways.EntityGateway) *CreateEntityUseCase {
	return &CreateEntityUseCase{gateway: gateway}
}

// Execute creates a new entity
func (uc *CreateEntityUseCase) Execute(entity entities.Entity) (entities.Entity, error) {
	return uc.gateway.Create(entity)
}
