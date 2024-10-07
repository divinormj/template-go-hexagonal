package postgresql

import (
	"template/internal/adapters/postgresql/presenters"

	"gorm.io/gorm"
)

// Migrate runs the migrations for the PostgreSQL database
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&presenters.UserModel{})
}
