package entities

import (
	"github.com/jinzhu/gorm"
)


type Post struct{
	gorm.Model
	Title string `gorm:"not null;unique;size:100" json:"title"`
	Slug  string `gorm:"not null;size:100" json:"slug"`
	Published bool `gorm:"not null;" json:"published"`
	Content string `gorm:"not null" json:"content"`
	Tags []*Tag `gorm:"many2many:post_tag;"`
}
