package dto

type Category struct {
	Name        string `json:"name"        validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=10"`
}

type CategoryUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
