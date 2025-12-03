package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null;type:text" json:"content"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
