package mongodb

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	presenters "template/internal/adapters/mongodb/presenters"
	"template/internal/domain/entities"
)

type EntityRepository struct {
	collection *mongo.Collection
}

// NewEntityRepository creates a new EntityRepository
func NewEntityRepository(client *mongo.Client, dbName string) *EntityRepository {
	collection := client.Database(dbName).Collection("entities")
	return &EntityRepository{collection: collection}
}

// Create inserts a new entity into the MongoDB collection
func (r *EntityRepository) Create(entity entities.Entity) (entities.Entity, error) {
	entityDTO := presenters.MapperEntityToDTO(entity)

	_, err := r.collection.InsertOne(context.Background(), entityDTO)
	if err != nil {
		return entities.Entity{}, err
	}

	return presenters.MapperDTOToEntity(entityDTO), nil
}

// Update modifies an existing entity in the MongoDB collection
func (r *EntityRepository) Update(entity entities.Entity) (entities.Entity, error) {
	filter := bson.M{"id": entity.ID}
	_, err := r.collection.UpdateOne(context.Background(), filter, bson.M{"$set": entity})
	if err != nil {
		return entities.Entity{}, err
	}
	return entity, nil
}

// Get retrieves an entity by ID from the MongoDB collection
func (r *EntityRepository) Get(id string) (entities.Entity, error) {
	var entityDTO presenters.EntityDTO
	objectID, errID := primitive.ObjectIDFromHex(id)
	if errID != nil {
		return entities.Entity{}, errors.New("invalid ID format")
	}
	filter := bson.M{"_id": objectID}
	err := r.collection.FindOne(context.Background(), filter).Decode(&entityDTO)
	if err != nil {
		return entities.Entity{}, err
	}
	return presenters.MapperDTOToEntity(entityDTO), nil
}

// Delete removes an entity by ID from the MongoDB collection
func (r *EntityRepository) Delete(id string) error {
	filter := bson.M{"id": id}
	_, err := r.collection.DeleteOne(context.Background(), filter)
	return err
}
