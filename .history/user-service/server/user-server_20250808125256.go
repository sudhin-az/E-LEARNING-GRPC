package server

import (
	"context"
	"fmt"
	"user-service/db"
	pb "user-service/proto"
)
type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo *db.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error){
	user, err := s.repo.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create User %v", err)
	}
	return &pb.UserResponse{
		Id: user.ID,
		Name: user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.UserID) (*pb.UserResponse, error) {
	users, err := s.repo.GetUser()
}