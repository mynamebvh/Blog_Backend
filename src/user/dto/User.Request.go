package dto

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   bool   `json:"gender"`
}

type UserUpdate struct {
	Fullname string `json:"fullname"`
	Gender   bool   `json:"gender"`
}