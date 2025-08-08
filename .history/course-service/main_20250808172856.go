package main

import (
	"course-service/server"
	"log"
	"net"
	pb "course-service/proto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterCourseServiceServer(s, &server.CourseServer{})

	log.Println("CourseService running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}