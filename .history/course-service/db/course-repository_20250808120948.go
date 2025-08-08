package db

import (
	"context"
	"fmt"

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
		return  nil, fmt.Errorf("failed to migrate models: %v", err)
	}

	return &CourseRepository{conn: conn}, nil
}

func (r *CourseRepository) CreateCourse(ctx context.Context, t) {
	
}