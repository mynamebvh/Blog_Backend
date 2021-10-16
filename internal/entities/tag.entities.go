package entities

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name        string    `gorm:"not null;unique;size:30" json:"name"`
	Description string    `gorm:"size:1000" json:"description"`
	Slug        string    `gorm:"not null;unique;size:100" json:"slug"`
	PostTag     []PostTag `gorm:"foreignKey:TagID" json:"post_tag"`
}
