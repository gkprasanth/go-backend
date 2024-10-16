package database

import "backend/models"

 

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}