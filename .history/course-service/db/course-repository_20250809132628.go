package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Course struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
}

type Enrollment struct {
	ID       string `gorm:"primaryKey"`
	UserID   string
	CourseID string
}

type CourseRepository struct {
	conn *gorm.DB
}

func NewCourseRepository(dbURL string) (*CourseRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}

	if err := conn.AutoMigrate(&Course{}, &Enrollment{}); err != nil {
		return nil, fmt.Errorf("failed to migrate models: %v", err)
	}

	return &CourseRepository{conn: conn}, nil
}

func (r *CourseRepository) CreateCourse(ctx context.Context, title, description string) (Course, error) {
	course := Course{
		ID:          uuid.New().String(),
		Title:       title,
		Description: description,
	}

	err := r.conn.Create(&course).Error
	if err != nil {
		return Course{}, fmt.Errorf("failed to create course: %v", err)
	}
	return course, nil
}

func (r *CourseRepository) GetAllCourses(ctx context.Context) ([]Course, error){
	var course []Course
	err := r.conn.Find(&course).Error
	if err != nil {
		return []Course{}, fmt.Errorf("failed to get courses %v", err)
	}
	fmt.Println("Courses Retrieved Successfully", )
	return  course, nil
}

func (r *CourseRepository) EnrollUser(ctx context.Context, userID, courseID string) error {
	enrollment := Enrollment{
		ID: uuid.NewString(),
		UserID: userID,
		CourseID: courseID,
	}

	if err := r.conn.Create(&enrollment).Error; err != nil {
		return fmt.Errorf("failed to enroll user: %v", err)
	}
	return  nil
}