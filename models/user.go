package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `gorm:"unique;not null" json:"username"`
	Password string    `gorm:"not null" json:"-"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Posts    []Post    `gorm:"foreignKey:UserID" json:"-"`
	Comments []Comment `gorm:"foreignKey:UserID" json:"-"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
