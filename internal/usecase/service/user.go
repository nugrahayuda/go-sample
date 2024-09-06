package service

import (
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
func (s *UserService) CreateUser(id string) (bool, error) {
	panic("unimplemented")
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
