package service

//go:generate mockgen -destination=mocks/user.go -source=user.go -package=service_mock

import (
	"context"
	"errors"
)

//go:generate mockgen -destination=mocks/user.go -source=user.go -package=service_mock

type UserService interface {
	GetById(ctx context.Context, id int32) (*User, error)
	GetByNickName(ctx context.Context, nickName string) (*User, error)
	List(ctx context.Context) ([]*User, error)
}

type User struct {
	Id       int32
	NickName string
}

var ErrNotFound = errors.New("not found")
