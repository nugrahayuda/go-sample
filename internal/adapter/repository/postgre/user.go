package repository

import (
	"database/sql"
	"fmt"
	"integrationtests/internal/domain/model"
	"integrationtests/internal/domain/repository"
)

func NewRepoUser(db *sql.DB) repository.UserRepository  {
	return &repoUser{db: db}
}

type repoUser struct {
	db *sql.DB
}

// Implements the RepoUser interface
func (r *repoUser) Create(userID string) (bool, error) {
	// Implementation
	return false, nil
}

func (r *repoUser) Delete(userID string) (bool, error) {
	// Implementation
	return false, nil
}

func (r *repoUser) GetUserByID(userID string) (model.UserData, error) {
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
