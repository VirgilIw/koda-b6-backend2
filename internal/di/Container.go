package di

import (
	"github.com/virgilIw/koda-b6-backend2/internal/handler"
	"github.com/virgilIw/koda-b6-backend2/internal/model"
	"github.com/virgilIw/koda-b6-backend2/internal/repository"
	"github.com/virgilIw/koda-b6-backend2/internal/service"
)

type Container struct {
	user        *[]model.UserModel
	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handler.UserHandler
}

func NewContainer() *Container {
	var DataUser []model.UserModel

	container := Container{
		user: &DataUser,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.user)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}
