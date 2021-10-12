package dto

type UserLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type UserRequest struct {
	Fullname string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
	Gender   bool   `validate:"required"`
}

type UserUpdate struct {
	Fullname string `validate:"required"`
	Gender   bool   `validate:"required"`
}
