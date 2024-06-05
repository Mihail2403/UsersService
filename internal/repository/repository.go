package repository

import (
	"github.com/jmoiron/sqlx"
	"go.mongodb.org/mongo-driver/mongo"
)

type Users interface {
	// GetAll() ([]entity.User, error)
	// GetByID(id int) (*entity.User, error)
	Create(user *UsersPostgresStruct) error
	// Update(user *entity.User) error
	// Delete(id int) error
	// GetByIDArray(ids []int) ([]entity.User, error)
}
type Img interface {
	// GetAll() ([]string, error)
	// GetByID(id string) (string, error)
	// GetByIDArray(id []string) ([]string, error)
	Create(img string) (string, error)
}

type Repository struct {
	Users
	Img
}

func NewRepository(db *sqlx.DB, mongoDB *mongo.Database) *Repository {
	return &Repository{
		Users: NewUsersRepo(db),
		Img:   NewImgRepo(mongoDB),
	}
}
