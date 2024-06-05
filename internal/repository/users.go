package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

type UsersPostgresStruct struct {
	ID     int    `db:"id"`
	Name   string `db:"name"`
	UserID int    `db:"user_id"`
	ImgID  string `db:"img_id"`
}

func (r *UsersRepo) Create(user *UsersPostgresStruct) error {
	query := fmt.Sprintf("INSERT INTO %s (name, user_id, img_id) VALUES ($1, $2, $3)", UsersTable)
	_, err := r.db.Exec(query, user.Name, user.UserID, user.ImgID)
	return err
}
