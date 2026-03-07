package model

import "time"

type UserDb struct {
	Id         int        `db:"id"`
	Fullname   string     `db:"fullname"`
	Email      string     `db:"email"`
	Password   string     `db:"password"`
	Picture    *string    `db:"picture"`
	Phone      *string    `db:"phone"`
	Address    *string    `db:"address"`
	Role       *string    `db:"role"`
	Created_At *time.Time `db:"created_at"`
}
