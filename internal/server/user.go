package server

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/EugeneImbro/chat-backend/internal/service"
)

func (s *Server) GetUserById(ctx context.Context, request *GetUserByIdRequest) (*User, error) {
	model, err := s.us.GetById(ctx, request.GetId())
	if err != nil {
		if errors.Is(err, service.ErrNotFound) {
			return nil, status.Errorf(codes.NotFound, fmt.Sprintf("user with id %d not found", request.GetId()))
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("cannot get user: %s", err.Error()))
	}
	return &User{Id: model.Id, NickName: model.NickName}, nil
}

func (s *Server) GetUserByNickName(ctx context.Context, request *GetUserByNickNameRequest) (*User, error) {
	model, err := s.us.GetByNickName(ctx, request.GetNickName())
	if err != nil {
		return nil, err
	}
	return &User{Id: model.Id, NickName: model.NickName}, nil
}

func (s *Server) GetUsers(ctx context.Context, empty *emptypb.Empty) (*GetUsersResponse, error) {
	models, err := s.us.List(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*User, len(models))
	for i, u := range models {
		users[i] = &User{Id: u.Id, NickName: u.NickName}
	}
	return &GetUsersResponse{Users: users}, nil
}
