package handlers

import (
	"context"

	pb "mithril-micro/usrv/pb"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.UserServiceServer {
	return userserviceService{}
}

type userserviceService struct{}

func (s userserviceService) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	var resp pb.CreateUserRes
	return &resp, nil
}

func (s userserviceService) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error) {
	var resp pb.User
	return &resp, nil
}

func (s userserviceService) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	var resp pb.UpdateUserRes
	return &resp, nil
}

func (s userserviceService) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	var resp pb.LoginRes
	return &resp, nil
}
