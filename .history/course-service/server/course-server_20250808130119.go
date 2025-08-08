package server

import "course-service/db"

type CourseServer struct {
	pb.UnimplementedCourseServiceServer
	repo *db.CourseRepository
}
