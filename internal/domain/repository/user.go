package repository

import "integrationtests/internal/domain/model"

type UserRepository interface {
	Create(userID string) (bool, error)
	Delete(userID string) (bool, error)
	GetUserByID(userID string) (model.UserData, error)
}