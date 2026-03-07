package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/di"
)

func InitRouter(app *gin.Engine, container *di.Container) {
	RouterUser(app, container.UserHandler())
}
