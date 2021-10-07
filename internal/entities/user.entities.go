package entities

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Fullname string `gorm:"not null;size:30" json:"fullname"`
	Email    string `gorm:"unique;not null;size:100" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Gender   bool `gorm:"type:bool;not null" json:"gender"`
	Post []Post `gorm:"foreignKey:id"`
}