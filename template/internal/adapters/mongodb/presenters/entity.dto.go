package mongodb

import "go.mongodb.org/mongo-driver/bson/primitive"

// EntityDTO represents the data transfer object for the entity
type EntityDTO struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
}
