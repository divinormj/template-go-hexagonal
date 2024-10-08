// Package presenters provides the definition of the DTO struct.
package presenters

// EntityDTO represents the data transfer object for the entity
type EntityDTO struct {
	Name string `json:"name" binding:"required"`
}
