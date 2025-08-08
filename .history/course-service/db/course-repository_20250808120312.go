package db

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
	conn *g
}
