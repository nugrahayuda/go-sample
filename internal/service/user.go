package service

import "integrationtests/internal/repo"


type userService struct {
	repo repo.RepoUser
}

func (s *userService) GetService (id uint32) (repo.UserData, error) {
	data, err := s.repo.GetUserByID(id)
    if err != nil {
        return data, err
    }

    return data, nil
}

func NewUserService(repo repo.RepoUser) *userService {
    return &userService{repo: repo}
}