// Package containers contains the dependency injection for the domain and entrypoints layer.
package containers

import (
	"template/internal/adapters/mongodb"
	"template/internal/adapters/postgresql"
	"template/internal/domain/usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// DomainContainer holds the application dependencies for the domain layer
type DomainContainer struct {
	CreateEntityUseCase *usecases.CreateEntityUseCase
	UpdateEntityUseCase *usecases.UpdateEntityUseCase
	GetEntityUseCase    *usecases.GetEntityUseCase
	DeleteEntityUseCase *usecases.DeleteEntityUseCase
	CreateUserUseCase   *usecases.CreateUserUseCase
	GetUserUseCase      *usecases.GetUserUseCase
}

// NewDomainContainer initializes the domain container
func NewDomainContainer(clientMongo *mongo.Client, clientPostgres *gorm.DB, dbName string) *DomainContainer {
	mongodb := mongodb.NewEntityRepository(clientMongo, dbName)
	postgres := postgresql.NewUserRepository(clientPostgres)

	return &DomainContainer{
		CreateEntityUseCase: usecases.NewCreateEntityUseCase(mongodb),
		UpdateEntityUseCase: usecases.NewUpdateEntityUseCase(mongodb),
		GetEntityUseCase:    usecases.NewGetEntityUseCase(mongodb),
		DeleteEntityUseCase: usecases.NewDeleteEntityUseCase(mongodb),
		CreateUserUseCase:   usecases.NewCreateUserUseCase(postgres),
		GetUserUseCase:      usecases.NewGetUserUseCase(postgres),
	}
}
