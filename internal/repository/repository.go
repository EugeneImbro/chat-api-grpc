package repository

import "github.com/jmoiron/sqlx"

type User interface {
}

type Repository struct {
	User
}

func NewRepository(r *sqlx.DB) *Repository {
	return &Repository{}
}
