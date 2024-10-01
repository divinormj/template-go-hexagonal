package usecases

import (
	"template/internal/domain/gateways"
)

// DeleteEntityUseCase defines the use case for deleting an entity
type DeleteEntityUseCase struct {
	gateway gateways.EntityGateway
}

// NewDeleteEntityUseCase creates a new delete entity use case
func NewDeleteEntityUseCase(gateway gateways.EntityGateway) *DeleteEntityUseCase {
	return &DeleteEntityUseCase{gateway: gateway}
}

// Execute deletes an entity by ID
func (uc *DeleteEntityUseCase) Execute(id string) error {
	return uc.gateway.Delete(id)
}
