package server

import (
	"context"
	"course-service/db"
	pb "course-service/proto"
)

type CourseServer struct {
	pb.UnimplementedCourseServiceServer
	repo *db.CourseRepository
}

func NewCourseServer(repo *db.CourseRepository) *CourseServer {
	return &CourseServer{repo: repo}
}

func (s *CourseServer) CreateCourse(ctx context.Context, req *pb.CourseRequest) (*pb.CourseResponse, error){
	err := s.repo.CreateCourse(ctx, r)
}
