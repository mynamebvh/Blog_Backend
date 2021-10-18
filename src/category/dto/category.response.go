package dto

import "time"

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perpage"`
	Total   int `json:"total"`
}

type CategoryRaw struct {
	Title       string    `json:"title"`
	PostID      uint      `json:"post_id"`
	PostSlug    string    `json:"post_slug"`
	FullName    string    `json:"fullname"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"create_at"`
}

type CategoryResponse struct {
	CategoryRaw []CategoryRaw
	Pagination  Pagination
}
