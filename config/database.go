package config

import (
	_ "github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	//DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	//DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	//DB, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	dsn := "root:123456@tcp(127.0.0.1:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	log.Println("Database connection established")
}
