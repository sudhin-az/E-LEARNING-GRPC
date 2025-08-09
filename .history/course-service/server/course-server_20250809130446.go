package server

import (
	"context"
	"github.com/sudhin-az/E-LEARNING-GRPC/course-service/db"
	pb "github.com/sudhin-az/E-LEARNING-GRPC/course-service/proto"
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

func (s *CourseServer) GetAllCourses(ctx context.Context, req *pb.Empty) (*pb.CourseList, error) {
	
}

func (s *CourseServer) EnrollCourse(ctx context.Context, req *pb.EnrollRequest) (*pb.EnrollResponse, error) {
	err := s.repo.EnrollUser(ctx, req.UserId, req.CourseId)
	if err != nil {
		return nil, fmt.Errorf("failed to Enroll user %v", err)
	}
	return &pb.EnrollResponse{Message: "Course Enrolled Successfully"}, nil
}

