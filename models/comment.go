package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	Content   string    `gorm:"not null;type:text" json:"content"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	PostID    uint      `json:"post_id"`
	Post      Post      `gorm:"foreignKey:PostID" json:"-"`
	CreatedAt time.Time `json:"created_at"`
}
