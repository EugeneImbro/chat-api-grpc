package server

import (
	"github.com/EugeneImbro/chat-backend/internal/service"
)

type Server struct {
	UnimplementedUserServiceServer
	services *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{services: s}
}
