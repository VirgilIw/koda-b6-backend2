package dto

type UserDataResponse struct {
	Id       int     `json:"id"`
	FullName string  `json:"fullname"`
	Email    string  `json:"email"`
	Phone    *string `json:"phone"`
}

type CreateUserResponse struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
