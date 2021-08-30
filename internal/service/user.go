package service

import (
	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository"
)

type UserService struct {
	repo repository.User
}

func (u *UserService) GetById(id int32) (*model.User, error) {
	return u.repo.GetById(id)
}

func (u *UserService) GetByNickName(nickName string) (*model.User, error) {
	return u.repo.GetByNickName(nickName)
}

func (u *UserService) GetAll() (*[]model.User, error) {
	return u.repo.GetAll()
}

func NewUserService(repos *repository.Repository) *UserService {
	return &UserService{repo: repos.User}
}
