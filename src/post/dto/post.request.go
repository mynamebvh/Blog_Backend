package dto

type Post struct {
	Title      string `json:"title"`
	Published  bool   `json:"published"`
	Content    string `json:"content"`
	UserID     string `json:"user_id"`
	CategoryID string `json:"category_id"`
}

type PostUpdate struct {
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
}
