// Package presenters contains the definition of the data models for the postgres database and the mapping to the corresponding domain entities
package presenters

import (
	"template/internal/domain/entities"
)

// UserModel is the model used for the PostgreSQL database
type UserModel struct {
	ID    string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name  string `gorm:"type:varchar(100);not null"`
	Email string `gorm:"type:varchar(100);not null"`
}

// TableName overrides the default table name for GORM
func (UserModel) TableName() string {
	return "users"
}

// ToDomain converts UserModel to domain Entity
func (model *UserModel) ToDomain() *entities.User {
	return &entities.User{
		ID:    model.ID,
		Name:  model.Name,
		Email: model.Email,
	}
}

// FromDomain converts domain User to UserModel
func FromDomain(entity *entities.User) *UserModel {
	return &UserModel{
		Name:  entity.Name,
		Email: entity.Email,
	}
}
