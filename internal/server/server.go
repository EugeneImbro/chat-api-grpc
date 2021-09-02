package server

import (
	"github.com/EugeneImbro/chat-backend/internal/service"
)

type Server struct {
	services *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{services: s}
}
