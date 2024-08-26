package service

import "integrationtests/internal/repo"

type userService struct {
	repo repo.RepoUser
}

type UserService interface {
	GetUserByID(id string) (repo.UserData, error)
	CreateUser(id string) (bool, error)
	DeleteUser(id string) (bool, error)
}

// CreateUser implements UserService.
func (s *userService) CreateUser(id string) (bool, error) {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (s *userService) DeleteUser(id string) (bool, error) {
	panic("unimplemented")
}

func (s *userService) GetUserByID(id string) (repo.UserData, error) {
	data, err := s.repo.GetUserByID(id)
	if err != nil {
		return data, err
	}

	return data, nil
}

func NewUserService(repo repo.RepoUser) UserService {
	return &userService{repo: repo}
}
