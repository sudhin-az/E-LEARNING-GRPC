package main

import (
	"log"
	"net"

	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/db"
	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/proto"
	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/server"

	"google.golang.org/grpc"
)

func main() {

	dbURL := "postgres://postgres:sudhin123@localhost:5432/e-learning-grpc"
	//Create DB connection
	userRepo, err := db.NewUserRepository(dbURL)

	//Create gRPC Sever
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, server.NewUserServer(userRepo))

	log.Println("UserService running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
