package models

import "gorm.io/gorm"

type Role string
const (
	Admin Role = "admin"
	NormalUser Role = "user"
)

// User represents a user in the system
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `json:"password"`
	Role     Role `json:"role"` // e.g., "admin" or "user"
}
