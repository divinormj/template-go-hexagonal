package usecases

import (
	"template/internal/domain/entities"
	"template/internal/domain/gateways"
)

// GetUserUseCase defines the use case for retrieving an entity
type GetUserUseCase struct {
	gateway gateways.UserGateway
}

// NewGetUserUseCase creates a new get user use case
func NewGetUserUseCase(gateway gateways.UserGateway) *GetUserUseCase {
	return &GetUserUseCase{gateway: gateway}
}

// Execute retrieves an entity by ID
func (usecase *GetUserUseCase) Execute(id string) (entities.User, error) {
	return usecase.gateway.Get(id)
}
