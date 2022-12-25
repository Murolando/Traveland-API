package service

import (
	"crypto/sha256"
	"fmt"
	"traveland/ent"
	"traveland/pkg/repository"
)

type UserCRUDService struct {
	repo repository.User
}

func NewUserCRUDService(repo repository.User) *UserCRUDService {
	return &UserCRUDService{
		repo: repo,
	}
}

func (s UserCRUDService) GetUserByID(id int) (ent.User, error) {
	return s.repo.GetUserByID(id)
}

func (s UserCRUDService) GetAllUsers() ([]ent.User, error) {
	return s.repo.GetAllUsers()
}
func (s UserCRUDService) GetUsersByRole(role_id int) ([]ent.User, error) {
	return s.repo.GetUsersByRole(role_id)
}
func (s UserCRUDService) UpdateUserInfo(user ent.User) (bool, error) {
	user.Password = s.generateHashPassword(user.Password)
	return s.repo.UpdateUserInfo(user)
}

func (s UserCRUDService) generateHashPassword(password string) string {

	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}
