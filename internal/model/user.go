package model

type UserModel struct {
	Id       int    `db:"id"`
	FullName string `db:"fullname"`
	Email    string `db:"email"`
}
