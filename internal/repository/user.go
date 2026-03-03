package repository

import (
	"fmt"

	"github.com/virgilIw/koda-b6-backend2/internal/model"
)

var userData = []model.UserModel{
	{
		Id:       1,
		FullName: "gilamb",
		Email:    "test1@gmail.com",
	},
}

type UserRepository struct {
	db []model.UserModel
}

func NewUserRepository(db *[]model.UserModel) *UserRepository {
	return &UserRepository{
		db: userData,
	}
}

func (r *UserRepository) GetUsers() []model.UserModel {
	return r.db
}

func (r *UserRepository) GetUsersById(id int) (model.UserModel, error) {
	for _, v := range r.db {
		if v.Id == id {
			return v, nil
		}
	}

	return model.UserModel{}, fmt.Errorf("user not found")
}

func (r *UserRepository) AddUser(user model.UserModel) model.UserModel {
	r.db = append(r.db, user)
	return user
}

func (r *UserRepository) UpdateByEmail(email string, updated model.UserModel) (model.UserModel, error) {
	for i, v := range r.db {
		if v.Email == email {

			// update field
			r.db[i].FullName = updated.FullName
			r.db[i].Email = updated.Email

			return r.db[i], nil
		}
	}

	return model.UserModel{}, fmt.Errorf("user not found")
}

func (r *UserRepository) DeleteUserById(id int) {
	for i, v := range r.db {
		if v.Id == id {
			r.db[i] = model.UserModel{}
			return
		}
	}
}
