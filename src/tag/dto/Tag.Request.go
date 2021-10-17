package dto

type Tag struct {
	Name        string `json:"name"        validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=10"`
}

type TagUpdate struct {
	Name        string `json:"name"        validate:"required,min=2"`
	Description string `json:"description" validate:"required,min=10"`
	Slug        string `json:"slug"`
}
