package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Fullname string `gorm:"unique_index;not null;size:30" json:"fullname"`
	Email    string `gorm:"unique_index;not null;size:100" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Gender bool `gorm: "not null" json:"gender"`
}