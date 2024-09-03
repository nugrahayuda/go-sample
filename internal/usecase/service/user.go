package service

import (
	"integrationtests/internal/domain/model"
	"integrationtests/internal/domain/repository"
)

//go:generate mockgen -source=user.go -debug

type UserService struct {
	repo repository.UserRepository
}

// CreateUser implements UserService.
func (s *UserService) CreateUser(id string) (bool, error) {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (s *UserService) DeleteUser(id string) (bool, error) {
	panic("unimplemented")
}

func (s *UserService) GetUserByID(id string) (model.UserData, error) {
	data, err := s.repo.GetUserByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}
