package repository

import (
	"context"

	"github.com/pkg/errors"

	"github.com/EugeneImbro/chat-backend/internal/model"
)

//go:generate mockgen -destination=mock/repository.go -source=repository.go  -package=repo_mock

var ErrNotFound = errors.New("not found")

type Repository interface {
	GetUserByID(ctx context.Context, id int32) (*model.User, error)
	GetUserByNickName(ctx context.Context, nickName string) (*model.User, error)
	UserList(ctx context.Context) ([]*model.User, error)
}
