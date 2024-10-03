// Package containers contains the dependency injection for the domain and entrypoints layer.
package containers

import (
	"template/internal/adapters/mongodb"
	"template/internal/domain/usecases"

	"go.mongodb.org/mongo-driver/mongo"
)

// DomainContainer holds the application dependencies for the domain layer
type DomainContainer struct {
	CreateEntityUseCase *usecases.CreateEntityUseCase
	UpdateEntityUseCase *usecases.UpdateEntityUseCase
	GetEntityUseCase    *usecases.GetEntityUseCase
	DeleteEntityUseCase *usecases.DeleteEntityUseCase
}

// NewDomainContainer initializes the domain container
func NewDomainContainer(client *mongo.Client, dbName string) *DomainContainer {
	repository := mongodb.NewEntityRepository(client, dbName)

	return &DomainContainer{
		CreateEntityUseCase: usecases.NewCreateEntityUseCase(repository),
		UpdateEntityUseCase: usecases.NewUpdateEntityUseCase(repository),
		GetEntityUseCase:    usecases.NewGetEntityUseCase(repository),
		DeleteEntityUseCase: usecases.NewDeleteEntityUseCase(repository),
	}
}
