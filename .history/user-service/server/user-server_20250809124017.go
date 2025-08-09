package server

import (
	"context"
	"fmt"

	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/db"
	pb "github.com/sudhin-az/E-LEARNING-GRPC/user-service/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo *db.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	user, err := s.repo.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create User %v", err)
	}
	fmt.Println("Saving to DB:", user.Name, user.Email)
	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context) (*[]pb.UserResponse, error) {
	_, err := s.repo.g(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to Get Users %v", err)
	}
	return &[]pb.UserResponse{}, nil
}
