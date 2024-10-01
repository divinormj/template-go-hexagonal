package mongodb

import (
	"template/internal/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MapEntityToDTO maps an Entity to an EntityDTO
func MapperEntityToDTO(entity entities.Entity) EntityDTO {
	return EntityDTO{
		ID:   primitive.NewObjectID(),
		Name: entity.Name,
	}
}

// MapDTOToEntity maps an EntityDTO to an Entity
func MapperDTOToEntity(dto EntityDTO) entities.Entity {
	return entities.Entity{
		ID:   dto.ID.Hex(),
		Name: dto.Name,
	}
}
