package containers

import (
	v1 "template/internal/adapters/entrypoints/v1"
)

// EntrypointContainer holds the application dependencies for the HTTP layer
type EntrypointContainer struct {
	V1EntityController *v1.EntityController
	V1UserController   *v1.UserController
}

// NewEntrypointContainer initializes the Entrypoint container
func NewEntrypointContainer(domainContainer *DomainContainer) *EntrypointContainer {
	return &EntrypointContainer{
		V1EntityController: v1.NewEntityController(
			domainContainer.CreateEntityUseCase,
			domainContainer.DeleteEntityUseCase,
			domainContainer.GetEntityUseCase,
			domainContainer.UpdateEntityUseCase,
		),
		V1UserController: v1.NewUserController(
			domainContainer.CreateUserUseCase,
			domainContainer.GetUserUseCase,
		),
	}
}
