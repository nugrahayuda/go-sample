package service

import (
	"context"
	"integrationtests/internal/adapter/repository/postgre/db"
	"integrationtests/internal/domain/model"
	"integrationtests/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepositoryInterface
}

func NewUserService(repo repository.UserRepositoryInterface) *UserService {
	return &UserService{repo: repo}
}

// CreateUser implements UserService.
func (s *UserService) CreateUser(user model.User) error {
	ctx := context.Background()
	tx := db.DBCon.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := s.repo.CreateUser(tx, user)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser implements UserService.
func (s *UserService) DeleteUser(id string) (bool, error) {
	panic("unimplemented")
}

func (s *UserService) GetUserByID(id string) (model.User, error) {
	data, err := s.repo.GetUserByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}
