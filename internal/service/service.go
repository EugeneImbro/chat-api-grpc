package service

import "github.com/EugeneImbro/chat-backend/internal/repository"

type User interface {
}

type Service struct {
	User
}

func NewService(r *repository.Repository) *Service {
	return &Service{}
}
