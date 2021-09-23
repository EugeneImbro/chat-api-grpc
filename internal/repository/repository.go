package repository

import (
	"context"

	"github.com/pkg/errors"
)

//go:generate mockgen -destination=mock/repository.go -source=repository.go  -package=repo_mock

var ErrNotFound = errors.New("not found")

type Repository interface {
	GetUserByID(ctx context.Context, id int32) (*User, error)
	GetUserByNickName(ctx context.Context, nickName string) (*User, error)
	UserList(ctx context.Context) ([]*User, error)
}

type User struct {
	Id       int32
	NickName string
}
