// Package entrypoints contains the controllers that handle HTTP requests and call the use cases
package entrypoints

import (
	"net/http"
	"template/internal/adapters/entrypoints/v1/presenters"
	"template/internal/domain/entities"
	"template/internal/domain/usecases"

	"github.com/gin-gonic/gin"
)

// EntityController handles the operations for the entity
type EntityController struct {
	createUseCase *usecases.CreateEntityUseCase
	deleteUseCase *usecases.DeleteEntityUseCase
	getUseCase    *usecases.GetEntityUseCase
	updateUseCase *usecases.UpdateEntityUseCase
}

// NewEntityController creates a new EntityController
func NewEntityController(
	createUseCase *usecases.CreateEntityUseCase,
	deleteUseCase *usecases.DeleteEntityUseCase,
	getUseCase *usecases.GetEntityUseCase,
	updateUseCase *usecases.UpdateEntityUseCase,
) *EntityController {
	return &EntityController{
		createUseCase: createUseCase,
		deleteUseCase: deleteUseCase,
		getUseCase:    getUseCase,
		updateUseCase: updateUseCase,
	}
}

// RegisterRoutes registers the routes for the entity
func (controller *EntityController) RegisterRoutes(r *gin.Engine) {
	const entityIDRoute = "/entities/:id"

	r.POST("/entities", controller.CreateEntity)
	r.DELETE(entityIDRoute, controller.DeleteEntity)
	r.GET(entityIDRoute, controller.GetEntity)
	r.PUT(entityIDRoute, controller.UpdateEntity)
}

// CreateEntity handles the creation of a new entity
// @Summary Create a new entity
// @Description Create a new entity with the provided details
// @Tags entities
// @Accept json
// @Produce json
// @Param entity body presenters.EntityDTO true "Entity details"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /entities [post]
func (controller *EntityController) CreateEntity(c *gin.Context) {
	var dto presenters.EntityDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := presenters.MapperDTOToEntity(dto)
	_, err := controller.createUseCase.Execute(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, nil)
}

// DeleteEntity handles the deletion of an entity by ID
// @Summary Delete an entity by ID
// @Description Delete an entity by its ID
// @Tags entities
// @Param id path string true "Entity ID"
// @Success 204
// @Failure 404
// @Router /entities/{id} [delete]
func (controller *EntityController) DeleteEntity(c *gin.Context) {
	id := c.Param("id")
	if err := controller.deleteUseCase.Execute(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// GetEntity handles the retrieval of an entity by ID
// @Summary Get an entity by ID
// @Description Retrieve an entity by its ID
// @Tags entities
// @Produce json
// @Param id path string true "Entity ID"
// @Success 200 {object} presenters.EntityDTO
// @Failure 404
// @Router /entities/{id} [get]
func (controller *EntityController) GetEntity(c *gin.Context) {
	id := c.Param("id")
	entity, err := controller.getUseCase.Execute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		return
	}
	c.JSON(http.StatusOK, presenters.MapperEntityToDTO(entity))
}

// UpdateEntity handles the update of an existing entity
// @Summary Update an existing entity
// @Description Update an existing entity with the provided details
// @Tags entities
// @Accept json
// @Produce json
// @Param id path string true "Entity ID"
// @Param entity body presenters.EntityDTO true "Entity details"
// @Success 200 {object} presenters.EntityDTO
// @Failure 400
// @Failure 500
// @Router /entities/{id} [put]
func (controller *EntityController) UpdateEntity(c *gin.Context) {
	var entity entities.Entity
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity.ID = c.Param("id")
	updatedEntity, err := controller.updateUseCase.Execute(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedEntity)
}
