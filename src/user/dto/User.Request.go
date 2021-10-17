package dto

type UserLogin struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UserRequest struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email"    validate:"required"`
	Password string `json:"password" validate:"required"`
	Gender   bool   `json:"gender"   validate:"required"`
}

type UserUpdate struct {
	Fullname string `json:"fullname" validate:"required,min=5,max=30"`
	Gender   *bool  `json:"gender"   validate:"required"`
}
