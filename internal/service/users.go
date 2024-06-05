package service

import (
	"users-service/entity"
	"users-service/internal/repository"
)

type UsersService struct {
	repo *repository.Repository
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) Create(user *entity.User) error {
	return nil
}
