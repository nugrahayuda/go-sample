package repository

import (
	"integrationtests/internal/domain/model"

	"gorm.io/gorm"
)

// If not working, run mockgen in the terminal
//
//go:generate mockgen -source=internal/domain/repository/user.go -destination=test/unit/mock/user.go -package=mock_repository_user
type UserRepositoryInterface interface {
	CreateUser(tx *gorm.DB, user model.User) error
	Delete(userID string) (bool, error)
	GetUserByID(userID string) (model.User, error)
}
