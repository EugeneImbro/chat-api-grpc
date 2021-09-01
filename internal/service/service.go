package service

import (
	"github.com/EugeneImbro/chat-backend/internal/model"
	"github.com/EugeneImbro/chat-backend/internal/repository"
)

////go:generate mockgen -destination=mocks/service.go -source=service.go -package=mock

type User interface {
	GetById(id int32) (*model.User, error)
	GetByNickName(nickName string) (*model.User, error)
	GetAll() ([]*model.User, error)
}

type Service struct {
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{User: NewUserService(r)}
}
