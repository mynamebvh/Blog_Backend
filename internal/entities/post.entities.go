package entities

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title     string `gorm:"not null;unique;size:100" json:"title"`
	Slug      string `gorm:"not null;size:100" json:"slug"`
	Published bool   `gorm:"not null;" json:"published"`
	Content   string `gorm:"not null;" json:"content"`
	UserID    uint   `gorm:"not null;column:user_id" json:"user_id"`
}

// Tags      []*Tag `gorm:"many2many:post_tag;"`
//;
