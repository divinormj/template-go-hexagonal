// Package gateways provides the definition of the Gateway interface.
package gateways

import "template/internal/domain/entities"

// UserGateway defines the interface for entity persistence
type UserGateway interface {
	Create(entity entities.User) (entities.User, error)
	Get(id string) (entities.User, error)
}
