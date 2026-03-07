package service

import (
	"context"
	"fmt"

	"github.com/matthewhartstonge/argon2"
	"github.com/virgilIw/koda-b6-backend2/internal/dto"
	"github.com/virgilIw/koda-b6-backend2/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetUsers(ctx context.Context) ([]dto.UserDataResponse, error) {
	datas, err := u.repo.GetUsers(ctx)

	if err != nil {
		return []dto.UserDataResponse{}, err
	}
	fmt.Println("ERROR:", err)
	var result []dto.UserDataResponse

	for _, v := range datas {

		result = append(result, dto.UserDataResponse{
			Id:       v.Id,
			FullName: v.Fullname,
			Email:    v.Email,
			Phone:    v.Phone,
		})
	}

	return result, nil
}

func (u *UserService) GetUserById(ctx context.Context, id int) (dto.UserDataResponse, error) {
	data, err := u.repo.GetUserById(ctx, id)
	var user dto.UserDataResponse

	user = dto.UserDataResponse{
		Id:       data.Id,
		FullName: data.Fullname,
		Email:    data.Email,
		Phone:    data.Phone,
	}

	if err != nil {
		return dto.UserDataResponse{}, err
	}

	return user, nil
}

func (u *UserService) CreateUser(ctx context.Context, fullname, email, password string) (dto.CreateUserResponse, error) {
	argon := argon2.DefaultConfig()

	hash, err := argon.HashEncoded([]byte(password))
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	data := u.repo.CreateUser(ctx, fullname, email, string(hash))

	user := dto.CreateUserResponse{
		Id:       data.Id,
		FullName: data.Fullname,
		Email:    data.Email,
	}

	return user, nil
}

func (u *UserService) DeleteUser(ctx context.Context, id int) error {
	if err := u.repo.DeleteUser(ctx, id); err != nil {
		return err
	}

	return nil
}
