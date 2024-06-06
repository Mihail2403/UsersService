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
	imgId, err := s.repo.Img.Create(user.ImgBase64)
	if err != nil {
		return err
	}
	pgUsr := &repository.UsersPostgresStruct{
		Name:   user.Name,
		UserID: user.UserId,
		ImgID:  imgId,
	}
	err = s.repo.Users.Create(pgUsr)
	return err
}
