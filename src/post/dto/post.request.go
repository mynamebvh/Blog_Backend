package dto

type Tag struct {
	Name string `json:"name"              validate:"required"`
	Slug string `json:"slug"`
}
type Post struct {
	Title      string `json:"title"       validate:"required,min=6"`
	Published  bool   `json:"published"   validate:"required,min=6"`
	Content    string `json:"content"     validate:"required,min=50"`
	UserID     string `json:"user_id"     validate:"required"`
	CategoryID string `json:"category_id" validate:"required"`
	Tags       []Tag  `json:"tags"        validate:"required"`
}

type PostUpdate struct {
	Title     string `json:"title"`
	Published bool   `json:"published"`
	Content   string `json:"content"`
}
