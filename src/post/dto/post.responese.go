package dto

import "time"

type PostRaw struct {
	Title     string
	Content   string
	Fullname  string
	Slug      string
	UserID    uint
	Published bool
	TagSlug   string
	UpdateAt  time.Time
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perpage"`
	Total   int `json:"total"`
}

type PostEntitiesRaw struct {
	PostID    uint      `json:"post_id"`
	Title     string    `json:"title"`
	Fullname  string    `json:"fullname"`
	Slug      string    `json:"slug"`
	UserID    uint      `json:"user_id"`
	TagSlug   string    `json:"tag"`
	CreatedAt time.Time `json:"created_at"`
}

type PostPagination struct {
	Posts      []PostResponse `json:"posts"`
	Pagination Pagination     `json:"pagination"`
}

type PostResponse struct {
	PostID    uint      `json:"post_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Fullname  string    `json:"fullname"`
	Slug      string    `json:"slug"`
	UserID    uint      `json:"user_id"`
	Published bool      `json:"published"`
	TagSlug   []string  `json:"tag"`
	UpdateAt  time.Time `json:"created_at"`
}
