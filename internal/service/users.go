package service

import "users-service/internal/repository"

type UsersService struct {
	repo *repository.Repository
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{repo: repo}
}
