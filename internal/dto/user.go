package dto

type UserDto struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

type UserAllDto struct {
	Id       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Image    string `json:"image"`
	Address  string `json:"address"`
}

type CreateUserRequestDto struct {
	FullName string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserRequestDto struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}
