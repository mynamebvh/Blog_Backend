package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"not null;unique;size:30" json:"name"`
	Description    string `gorm:"not null;size:1000" json:"description"`
	slug string `gorm:"not null;unique;size:100" json:"slug"`
}