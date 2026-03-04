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

	product        *[]model.ProductModel
	productRepo    *repository.ProductRepository
	productService *service.ProductService
	productHandler *handler.ProductHandler
}

func NewContainer() *Container {
	var DataUser []model.UserModel
	var DataProduct []model.ProductModel

	container := Container{
		user:    &DataUser,
		product: &DataProduct,
	}

	container.initDependencies()

	return &container
}

func (c *Container) initDependencies() {
	// USER
	c.userRepo = repository.NewUserRepository(c.user)
	c.userService = service.NewUserService(c.userRepo)
	c.userHandler = handler.NewUserHandler(c.userService)

	// PRODUCT
	c.productRepo = repository.NewProductRepository(c.product)
	c.productService = service.NewProductService(c.productRepo)
	c.productHandler = handler.NewProductHandler(c.productService)
}

func (c *Container) UserHandler() *handler.UserHandler {
	return c.userHandler
}

func (c *Container) ProductHandler() *handler.ProductHandler {
	return c.productHandler
}
