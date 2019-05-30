package db

import "go-api-template/models"

func StartMigration() {
	db := Get()
	db.AutoMigrate(&models.User{})
}
