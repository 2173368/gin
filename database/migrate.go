package database

import (
	"log"

	"gin/config"
	"gin/models"
)

func Migrate() {
	// 先禁用外键检查
	config.DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	// 按正确顺序迁移表（先创建被引用的表）
	err := config.DB.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Comment{},
	)

	// 重新启用外键检查
	config.DB.Exec("SET FOREIGN_KEY_CHECKS = 1")

	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	log.Println("Database migration completed")
}
