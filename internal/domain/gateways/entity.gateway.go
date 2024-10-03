// Package gateways provides the definition of the Gateway interface.
package gateways

import "template/internal/domain/entities"

// EntityGateway defines the interface for entity persistence
type EntityGateway interface {
	Create(entity entities.Entity) (entities.Entity, error)
	Update(entity entities.Entity) (entities.Entity, error)
	Get(id string) (entities.Entity, error)
	Delete(id string) error
}
