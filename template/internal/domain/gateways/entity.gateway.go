package gateways

import "template/internal/domain/entities"

// ItemGateway defines the interface for entity persistence
type EntityGateway interface {
	Create(entity entities.Entity) (entities.Entity, error)
	Update(entity entities.Entity) (entities.Entity, error)
	Get(id string) (entities.Entity, error)
	Delete(id string) error
}
