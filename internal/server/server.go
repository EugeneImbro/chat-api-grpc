package server

import (
	"github.com/EugeneImbro/chat-backend/internal/service"
)
type Server struct {
	us service.UserService
}

func NewServer(s service.UserService) *Server {
	return &Server{us: s}
}
