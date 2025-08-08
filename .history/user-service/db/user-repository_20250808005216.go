package db

import "gorm.io/gorm"

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}

type UserRepository struct {
	conn *gorm.DB
}

func ()  {
	
}