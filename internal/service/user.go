package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository"
)

//go:generate mockgen -destination=mocks/user.go -source=user.go -package=mock

type User interface {
	GetById(ctx context.Context, id int32) (*model.User, error)
	GetByNickName(ctx context.Context, nickName string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}

type UserService struct {
	repo repository.Repository
}

func NewUserService(r repository.Repository) *UserService {
	return &UserService{repo: r}
}

var ErrNotFound = errors.New("not found")

func (u *UserService) GetById(ctx context.Context, id int32) (*model.User, error) {
	usr, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("cannot get user from repository: %w", err)
	}
	return usr, err
}

func (u *UserService) GetByNickName(ctx context.Context, nickName string) (*model.User, error) {
	return u.repo.GetUserByNickName(ctx, nickName)
}

func (u *UserService) List(ctx context.Context) ([]*model.User, error) {
	return u.repo.UserList(ctx)
}
