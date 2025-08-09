package main

import (
	"log"
	"net"

	"github.com/sudhin-az/E-LEARNING-GRPC/course-service/db"
	"github.com/sudhin-az/E-LEARNING-GRPC/course-service/proto"
	"github.com/sudhin-az/E-LEARNING-GRPC/course-service/server"

	"google.golang.org/grpc"
)

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/e-learning-grpc?sslmode=disable"
	courseRepo, err := db.NewCourseRepository(dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterCourseServiceServer(s, server.NewCourseServer(courseRepo))

	log.Println("CourseService running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
