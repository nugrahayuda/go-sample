package repository

import (
	"fmt"
	"integrationtests/internal/domain/model"
	"integrationtests/internal/domain/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepositoryInterface {
	return &userRepository{db: db}
}

// Implements the RepoUser interface
func (r *userRepository) CreateUser(tx *gorm.DB, user model.User) error {
	var ur model.User

	err := tx.Create(&ur).Error
	return err
}

func (r *userRepository) Delete(userID string) (bool, error) {
	// Implementation
	return false, nil
}

func (r *userRepository) GetUserByID(userID string) (model.User, error) {
	// Create an instance to hold the result
	var user model.User

	// Use GORM to retrieve the user by ID
	err := r.db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, fmt.Errorf("user not found")
		}
		return user, err
	}

	return user, nil
}
