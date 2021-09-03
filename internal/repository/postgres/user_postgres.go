package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) GetById(ctx context.Context, id int32) (*model.User, error) {
	var user model.User

	if err := r.db.GetContext(ctx, &user,
		"SELECT * FROM users WHERE id=$1",
		id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (r *UserPostgres) GetByNickName(ctx context.Context, nickName string) (*model.User, error) {
	var user model.User
	if err := r.db.GetContext(ctx,
		&user,
		"SELECT * FROM users WHERE nickname=$1",
		nickName); err != nil {
		return nil, fmt.Errorf("failed to get user by nickname: %w", err)
	}
	return &user, nil
}

func (r *UserPostgres) List(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := r.db.SelectContext(ctx, users,
		"SELECT * FROM users",
	); err != nil {
		return nil, fmt.Errorf("failed to get user list: %w", err)
	}
	return users, nil
}
