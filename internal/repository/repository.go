package repository

import "github.com/jmoiron/sqlx"

type Users interface {
	// GetAll() ([]entity.User, error)
	// GetByID(id int) (*entity.User, error)
	// Create(user *entity.User) error
	// Update(user *entity.User) error
	// Delete(id int) error
	// GetByIDArray(ids []int) ([]entity.User, error)
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
	}
}