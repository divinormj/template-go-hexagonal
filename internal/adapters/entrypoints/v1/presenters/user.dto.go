// Package presenters provides the definition of the DTO struct.
package presenters

import "template/internal/domain/entities"

// UserDTO represents the data transfer object for the user
type UserDTO struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

// MapperDTOToUser converts UserDTO to domain Entity
func MapperDTOToUser(dto UserDTO) entities.User {
	return entities.User{
		Name:  dto.Name,
		Email: dto.Email,
	}
}

// MapperUserToDTO converts domain User to UserDTO
func MapperUserToDTO(entity entities.User) UserDTO {
	return UserDTO{
		Name:  entity.Name,
		Email: entity.Email,
	}
}
