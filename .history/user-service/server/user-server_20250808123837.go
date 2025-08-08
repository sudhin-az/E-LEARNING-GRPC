package server

import "user-service/db"

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.
}