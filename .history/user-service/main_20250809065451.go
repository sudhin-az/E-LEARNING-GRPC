package main

import (
	"log"
	"net"

	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/proto"
	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/server"

	"google.golang.org/grpc"
)

func main() {

	
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &server.UserServer{})

	log.Println("UserService running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
