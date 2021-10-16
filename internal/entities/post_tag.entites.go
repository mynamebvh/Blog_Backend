package entities

import "github.com/jinzhu/gorm"

type PostTag struct {
	gorm.Model
	PostID uint `gorm:"not null;column:post_id" json:"post_id"`
	TagID  uint `gorm:"not null;column:tag_id" json:"tag_id"`
}
