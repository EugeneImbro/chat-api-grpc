package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/EugeneImbro/chat-backend/internal/repository"
)

type pg struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) repository.Repository {
	return &pg{db: db}
}

func (r pg) GetUserByID(ctx context.Context, id int32) (*repository.User, error) {
	var user repository.User
	if err := r.db.GetContext(ctx, &user,
		"SELECT id, nickname FROM users WHERE id=$1",
		id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by id: %w", err)
	}
	return &user, nil
}

func (r pg) GetUserByNickName(ctx context.Context, nickName string) (*repository.User, error) {
	var user repository.User
	if err := r.db.GetContext(ctx,
		&user,
		"SELECT id, nickname FROM users WHERE nickname=$1",
		nickName); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("failed to get user by nickname: %w", err)
	}
	return &user, nil
}

func (r pg) UserList(ctx context.Context) ([]*repository.User, error) {
	var users []*repository.User
	if err := r.db.SelectContext(ctx, users,
		"SELECT id, nickname FROM users",
	); err != nil {
		return nil, fmt.Errorf("failed to get user list: %w", err)
	}
	return users, nil
}
