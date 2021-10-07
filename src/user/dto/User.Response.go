package dto

type UserResponse struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Gender   bool   `json:"gender"`
}

type JwtResponse struct {
	Type  string `json:"type"`
	Token string `json:"token"`
}