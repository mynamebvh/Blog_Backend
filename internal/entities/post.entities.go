package entities

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Title      string    `gorm:"not null;unique;size:100" json:"title"`
	Slug       string    `gorm:"not null;unique;size:100" json:"slug"`
	Published  bool      `gorm:"not null;" json:"published"`
	Content    string    `gorm:"not null;" json:"content"`
	UserID     uint      `gorm:"not null;column:user_id" json:"user_id"`
	CategoryID uint      `gorm:"not null;column:category_id" json:"category_id"`
	PostTag    []PostTag `gorm:"foreignKey:PostID" json:"post_tag"`
}
