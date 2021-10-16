package dto

type Tag struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
type Post struct {
	Title      string `json:"title"`
	Published  bool   `json:"published"`
	Content    string `json:"content"`
	UserID     string `json:"user_id"`
	CategoryID string `json:"category_id"`
	Tags       []Tag  `json:"tags"`
}

type PostUpdate struct {
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
}
