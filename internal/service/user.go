package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository"
)

type UserService struct {
	repo repository.User
}

func (u *UserService) GetById(ctx context.Context, id int32) (*model.User, error) {
	usr, err := u.repo.GetById(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("cannot get user from repository: %w", err)
	}
	return usr, err
}

func (u *UserService) GetByNickName(ctx context.Context, nickName string) (*model.User, error) {
	return u.repo.GetByNickName(ctx, nickName)
}

func (u *UserService) List(ctx context.Context) ([]*model.User, error) {
	return u.repo.List(ctx)
}

func NewUserService(repos *repository.Repository) *UserService {
	return &UserService{repo: repos.User}
}
