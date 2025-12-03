package database

import (
	"log"

	"gin/config"
	"gin/models"
)

func Migrate() {
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}
	
	log.Println("Database migration completed")
}