package service

import (
	"crypto/sha256"
	"fmt"
	"traveland/ent"
	"traveland/pkg/repository"
)
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}	

func (s AuthService) CreateUser(user ent.User) (int, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s AuthService) generateHashPassword(password string) string{
	
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x",hash)
}
