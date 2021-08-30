package repository

import (
	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/jmoiron/sqlx"
)

type User interface {
	GetById(id int32) (*model.User, error)
	GetByNickName(nickName string) (*model.User, error)
	GetAll() (*[]model.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
