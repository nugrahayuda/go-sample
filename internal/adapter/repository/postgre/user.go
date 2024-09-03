package repository

import (
	"database/sql"
	"fmt"
	"integrationtests/internal/domain/model"
	"integrationtests/internal/domain/repository"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepositoryInterface {
	return &userRepository{db: db}
}

// Implements the RepoUser interface
func (r *userRepository) Create(userID string) (bool, error) {
	// Implementation
	return false, nil
}

func (r *userRepository) Delete(userID string) (bool, error) {
	// Implementation
	return false, nil
}

func (r *userRepository) GetUserByID(userID string) (model.UserData, error) {
	// Create query to get user by ID
	var user model.UserData

	query := "SELECT id, name, role_id, is_active, birthday, phone_number FROM user WHERE id = ?"
	err := r.db.QueryRow(query, userID).Scan(&user.Id, &user.Name, &user.Role, &user.Status, &user.Birthday, &user.PhoneNumber)
	if err == sql.ErrNoRows {
		return user, fmt.Errorf("user not found")
	} else if err != nil {
		return user, err
	}

	return user, nil
}
