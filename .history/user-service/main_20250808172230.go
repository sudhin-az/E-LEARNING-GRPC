package main

import (
	"log"
	"net"
	"user-service/server"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "5000")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server.UserServer{})

	log.Println("UserService running on port 5000")
	if err := s.Serve(lis); err != nil {
		l
	}
}