package repository

import "integrationtests/internal/domain/model"

// If not working, run mockgen in the terminal
//
//go:generate mockgen -source=internal/domain/repository/user.go -destination=test/unit/mock/user.go -package=mock_repository_user
type UserRepositoryInterface interface {
	Create(userID string) (bool, error)
	Delete(userID string) (bool, error)
	GetUserByID(userID string) (model.UserData, error)
}
