package server

import (
	"context"
)

func (s *Server) GetUserById(ctx context.Context, request *GetUserByIdRequest) (*User, error) {
	model, err := s.services.User.GetById(request.GetId())
	if err != nil {
		return nil, err
	}
	return &User{Id: model.Id, NickName: model.NickName}, nil
}

func (s *Server) GetUserByNickName(ctx context.Context, request *GetUserByNickNameRequest) (*User, error) {
	model, err := s.services.User.GetByNickName(request.GetNickName())
	if err != nil {
		return nil, err
	}
	return &User{Id: model.Id, NickName: model.NickName}, nil
}

func (s *Server) GetUsers(ctx context.Context, request *GetUsersRequest) (*GetUsersResponse, error) {
	models, err := s.services.User.GetAll()
	if err != nil {
		return nil, err
	}
	users := make([]*User, len(models))
	for i, u := range models {
		users[i] = &User{Id: u.Id, NickName: u.NickName}
	}
	return &GetUsersResponse{Users: users}, nil
}
