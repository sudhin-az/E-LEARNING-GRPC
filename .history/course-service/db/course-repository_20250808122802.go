package db

import (
	"context"
	"fmt"

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

func (r *CourseRepository) GetCourse(ctx context.Context, id string) (Course, error){
	var course Course

	err := r.conn.Where("id = ?", id).First(&course).Error
	if err != nil {
		return  Course{}, fmt.Errorf("course not found: %v", err)
	}
	return  course, nil
}

func (r *CourseRepository) EnrollUser(ctx context.Context, userID, courseID string) error {
	enrollm
}