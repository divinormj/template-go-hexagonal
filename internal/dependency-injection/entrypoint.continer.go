package containers

import (
	v1 "template/internal/adapters/entrypoints/v1"
)

// EntrypointContainer holds the application dependencies for the HTTP layer
type EntrypointContainer struct {
	V1EntityController *v1.EntityController
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
	}
}
