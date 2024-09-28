package models

import "gorm.io/gorm"

// User defines the structure for the "users" table
type User struct {
	gorm.Model        // Adds fields ID, CreatedAt, UpdatedAt, DeletedAt
	Username   string `gorm:"type:varchar(100);unique" json:"username"`
	Email      string `gorm:"type:varchar(100);unique" json:"email"`
	Password   string `json:"-"`
}
