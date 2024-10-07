// Package usecases contains the application use cases.
package usecases

import (
	"template/internal/domain/entities"
	"template/internal/domain/gateways"
)

// CreateUserUseCase defines the use case for creating an entity
type CreateUserUseCase struct {
	gateway gateways.UserGateway
}

// NewCreateUserUseCase creates a new create user use case
func NewCreateUserUseCase(gateway gateways.UserGateway) *CreateUserUseCase {
	return &CreateUserUseCase{gateway: gateway}
}

// Execute creates a new user
func (createUser *CreateUserUseCase) Execute(user entities.User) (entities.User, error) {
	return createUser.gateway.Create(user)
}
