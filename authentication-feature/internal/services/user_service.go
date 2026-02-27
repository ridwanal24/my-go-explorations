package services

import (
	"authentication-feature/internal/models"
	"authentication-feature/internal/repositories"
)

type UserService interface {
	GetUser(id int) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetUser(id int) (*models.User, error) {
	return s.repo.FindById(id)
}
