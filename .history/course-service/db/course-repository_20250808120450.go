package db

import (
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
	conn, err := gorm.Open(postgres.)
}