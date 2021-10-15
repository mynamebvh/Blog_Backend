package dto

type Post struct {
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
}

type PostUpdate struct {
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
}
