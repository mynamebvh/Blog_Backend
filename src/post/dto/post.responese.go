package dto

import "time"

type PostRaw struct {
	Title     string
	Content   string
	Fullname  string
	Slug      string
	UserID    uint
	Published bool
	UpdateAt  time.Time
}

type PostResponse struct {
	Title     string
	Content   string
	Fullname  string
	Slug      []string
	UserID    uint
	Published bool
	UpdateAt  time.Time
}
