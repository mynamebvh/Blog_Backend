package dto

type Category struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CategoryUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
}
