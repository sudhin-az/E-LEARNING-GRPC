 package db

type Course struct {
	ID string `gorm:"primaryKey"`
	Title string 
	Description string
}

type Enrollment struct {
	I
}
