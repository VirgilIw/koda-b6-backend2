package service

import (
	"github.com/virgilIw/koda-b6-backend2/internal/dto"
	"github.com/virgilIw/koda-b6-backend2/internal/model"
	"github.com/virgilIw/koda-b6-backend2/internal/repository"
)

type UserService struct {
	repo   *repository.UserRepository
	lastID int
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo:   repo,
		lastID: 1,
	}
}

func (u *UserService) GetUsers() []dto.UserDto {
	data := u.repo.GetUsers()

	var result []dto.UserDto

	for _, v := range data {
		result = append(result, dto.UserDto{
			Id:       v.Id,
			FullName: v.FullName,
			Email:    v.Email,
		})
	}

	return result
}

func (u *UserService) GetUserById(id int) (dto.UserDto, error) {
	user, err := u.repo.GetUsersById(id)
	if err != nil {
		return dto.UserDto{}, err
	}

	return dto.UserDto{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}

func (u *UserService) CreateUser(req dto.CreateUserRequestDto) dto.UserDto {
	users := u.repo.GetUsers()

	newID := len(users) + 1

	newUser := model.UserModel{
		Id:       newID,
		FullName: req.FullName,
		Email:    req.Email,
	}

	u.repo.AddUser(newUser)

	return dto.UserDto{
		Id:       newUser.Id,
		FullName: newUser.FullName,
		Email:    newUser.Email,
	}
}

func (u *UserService) UpdateByEmail(email string, req dto.UpdateUserRequestDto) (dto.UserDto, error) {

	updated := model.UserModel{
		FullName: req.FullName,
		Email:    req.Email, // ini email baru
	}

	user, err := u.repo.UpdateByEmail(email, updated) // email lama untuk cari
	if err != nil {
		return dto.UserDto{}, err
	}

	return dto.UserDto{
		Id:       user.Id,
		FullName: user.FullName,
		Email:    user.Email,
	}, nil
}

func (u *UserService) DeleteUserById(id int) {
	u.repo.DeleteUserById(id)
}
