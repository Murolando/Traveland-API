package service

import (
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

func (s UserCRUDService) GetAllUsers() ([]ent.User,error) {
	return s.repo.GetAllUsers()
}
func (s UserCRUDService) GetUsersByRole(role_id int) ([]ent.User,error){
	return s.repo.GetUsersByRole(role_id)
}