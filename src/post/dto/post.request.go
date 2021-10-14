package dto

type Post struct {
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
	UserID    uint   `json:"user_id"`
}

type PostUpdate struct {
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
	Slug      string `json:"slug"`
}
