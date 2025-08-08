package server

import (
	"context"
	"course-service/db"
	pb "course-service/proto"
	"fmt"
)

type CourseServer struct {
	pb.UnimplementedCourseServiceServer
	repo *db.CourseRepository
}

func NewCourseServer(repo *db.CourseRepository) *CourseServer {
	return &CourseServer{repo: repo}
}

func (s *CourseServer) CreateCourse(ctx context.Context, req *pb.CourseRequest) (*pb.CourseResponse, error){
	course, err := s.repo.CreateCourse(ctx, req.Title, req.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to create Course %v", err)
	}
	return &pb.CourseResponse{
		Id: course.ID,
		Title: course.Title,
		Description: course.Description,
	}, nil
}

func (s *CourseServer) GetCourse(ctx context.Context, req *pb.CourseID) (*pb.CourseResponse, error) {
	course, err := s.repo.GetCourse(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to Get Courses %v", err)
	}
	return &pb.CourseResponse{
		Id: course.ID,
		Title: course.Title,
		Description: course.Description,
	}, nil
}

func (s *CourseServer) EnrollCourse(ctx context.Context, req *pb.EnrollRequest) (*pb.EnrollResponse, error) {
	msg, _ := s.repo.EnrollUser(ctx, req.UserId, req.CourseId)
	if err != nil {
		return nil, fmt.Errorf("failed to Enroll user %v", err)
	}
	return &pb.EnrollResponse{}, nil
}
