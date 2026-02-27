package services

import "authentication-feature/internal/repositories"

type AuthService interface {
	DoLogin(email, password string) bool
}

type authService struct {
	userRepo repositories.UserRepository
}

func NewAuthService(r repositories.UserRepository) AuthService {
	return &authService{userRepo: r}
}

func (s *authService) DoLogin(email, password string) bool {
	u, _ := s.userRepo.FindByEmail(email)

	if password != u.Password {
		return false
	}

	return true
}
