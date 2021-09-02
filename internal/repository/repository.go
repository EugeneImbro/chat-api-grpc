package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/EugeneImbro/chat-backend/internal/model"
)

//go:generate mockgen -destination=mock/repository.go -source=repository.go  -package=repo_mock

type User interface {
	GetById(id int32) (*model.User, error)
	GetByNickName(nickName string) (*model.User, error)
	GetAll() ([]*model.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserPostgres(db),
	}
}
