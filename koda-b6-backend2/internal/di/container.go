package di

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"github.com/virgilIw/koda-b6-backend2/internal/handler"
	"github.com/virgilIw/koda-b6-backend2/internal/repository"
	"github.com/virgilIw/koda-b6-backend2/internal/service"
)

type Container struct {
	db  *pgxpool.Pool
	rdb *redis.Client

	userRepo    *repository.UserRepository
	userService *service.UserService
	userHandler *handler.UserHandler
}

func NewContainer(db *pgxpool.Pool, rdb *redis.Client) *Container {
	container := &Container{
		db:  db,
		rdb: rdb,
	}

	container.initDependencies()

	return container
}

func (c *Container) initDependencies() {
	c.userRepo = repository.NewUserRepository(c.db, c.rdb)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)

}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}
