package server

import (
	pb "user-service/proto"
	"user-service/db"
)
type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo *db.UserRepository) *UserServer {
	return us
}