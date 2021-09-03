package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository/postgres"
)

//go:generate mockgen -destination=mock/repository.go -source=repository.go  -package=repo_mock

type User interface {
	GetById(ctx context.Context, id int32) (*model.User, error)
	GetByNickName(ctx context.Context, nickName string) (*model.User, error)
	List(ctx context.Context, ) ([]*model.User, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: postgres.NewUserPostgres(db),
	}
}
