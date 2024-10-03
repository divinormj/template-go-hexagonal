// Package presenters contains the DTOs (Data Transfer Objects) and mappers to transfer data between the domain layer and the MongoDB repository implementation.
package presenters

import "go.mongodb.org/mongo-driver/bson/primitive"

// EntityDTO represents the data transfer object for the persistence entity
type EntityDTO struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}
