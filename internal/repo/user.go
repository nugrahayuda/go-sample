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
	id          uint32
	name        string
	role        string
	status      string
	birthday    time.Time
	phoneNumber string
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

	query := "SELECT name, role, status, birthday, phone_number FROM users WHERE id = ?"
	err := r.db.QueryRow(query, userID).Scan(&user.id, &user.name, &user.role, &user.status, &user.birthday, &user.phoneNumber)
	if err == sql.ErrNoRows {
		return user, fmt.Errorf("user not found")
	} else if err != nil {
		return user, err
	}

	return user, nil
}
