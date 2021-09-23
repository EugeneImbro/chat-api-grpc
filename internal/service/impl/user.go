package impl

import (
	"context"
	"errors"
	"fmt"

	"github.com/EugeneImbro/chat-backend/internal/repository"
	"github.com/EugeneImbro/chat-backend/internal/service"
)

type us struct {
	repo repository.Repository
}

func NewUserService(r repository.Repository) service.UserService {
	return &us{repo: r}
}

func (u *us) GetById(ctx context.Context, id int32) (*service.User, error) {
	usr, err := u.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, fmt.Errorf("failed to GetByID: %w", err)
	}
	return (*service.User)(usr), err
}

func (u *us) GetByNickName(ctx context.Context, nickName string) (*service.User, error) {
	usr, err := u.repo.GetUserByNickName(ctx, nickName)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, service.ErrNotFound
		}
		return nil, fmt.Errorf("failed to GetByID: %w", err)
	}
	return (*service.User)(usr), err
}

func (u *us) List(ctx context.Context) ([]*service.User, error) {
	users, err := u.repo.UserList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to List: %w", err)
	}
	out := make([]*service.User, len(users))
	for i, v := range users {
		out[i] = (*service.User)(v)
	}
	return out, nil
}
