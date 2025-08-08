package main

import (
	"log"
	"net"
	pb "user-service/proto"
	"user-service/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "500")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server.UserServer{})

	log.Println("UserService running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
