package usecase

import (
	"pizza-tracker-go/internal/domain/entity"
	"pizza-tracker-go/internal/domain/repository"
)

type AuthService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Authenticate(username, password string) (*entity.User, error) {
	return s.userRepo.Authenticate(username, password)
}

func (s *AuthService) GetUserByID(id string) (*entity.User, error) {
	return s.userRepo.GetByID(id)
}
