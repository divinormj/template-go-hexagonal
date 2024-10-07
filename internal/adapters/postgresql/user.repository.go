// Package postgresql provides the implementation of the Gateways interface using PostgreSQL.
package postgresql

import (
	"errors"
	"template/internal/adapters/postgresql/presenters"
	"template/internal/domain/entities"
	"template/internal/domain/gateways"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository cria um novo repositório de usuários
func NewUserRepository(db *gorm.DB) gateways.UserGateway {
	return &UserRepository{db: db}
}

// Create insere um novo usuário no banco de dados
func (r *UserRepository) Create(entity entities.User) (entities.User, error) {
	userModel := presenters.FromDomain(&entity)

	if err := r.db.Create(userModel).Error; err != nil {
		return entities.User{}, err
	}

	return *userModel.ToDomain(), nil
}

// Get busca um usuário pelo ID no banco de dados
func (r *UserRepository) Get(id string) (entities.User, error) {
	var userModel presenters.UserModel

	if err := r.db.First(&userModel, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.User{}, errors.New("user not found")
		}
		return entities.User{}, err
	}

	return *userModel.ToDomain(), nil
}
