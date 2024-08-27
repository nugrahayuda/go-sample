package repo

import (
	"database/sql"
	"fmt"
	"time"
)

type RepoUser interface {
	Create(userID string) (bool, error)
	Delete(userID string) (bool, error)
	GetUserByID(userID string) (UserData, error)
}

type repoUser struct {
	db *sql.DB
}

type UserData struct {
	Id          uint32    `json:"id"`
	Name        string    `json:"name"`
	Role        string    `json:"role"`
	Status      string    `json:"status"`
	Birthday    time.Time `json:"birthday"`
	PhoneNumber string    `json:"phoneNumber"`
}

func NewRepoUser(db *sql.DB) RepoUser {
	return &repoUser{db: db}
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

func (r *repoUser) GetUserByID(userID string) (UserData, error) {
	// Create query to get user by ID
	var user UserData

	query := "SELECT id, name, role_id, is_active, birthday, phone_number FROM user WHERE id = ?"
	err := r.db.QueryRow(query, userID).Scan(&user.Id, &user.Name, &user.Role, &user.Status, &user.Birthday, &user.PhoneNumber)
	if err == sql.ErrNoRows {
		return user, fmt.Errorf("user not found")
	} else if err != nil {
		return user, err
	}

	return user, nil
}
