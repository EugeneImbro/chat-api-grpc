package service

import (
	"context"

	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository"
)

type UserService struct {
	repo repository.User
}

func (u *UserService) GetById(ctx context.Context, id int32) (*model.User, error) {
	return u.repo.GetById(ctx, id)
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
