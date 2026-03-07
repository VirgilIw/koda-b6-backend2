package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/handler"
)

func RouterUser(app *gin.Engine, userHandler *handler.UserHandler) {
	user := app.Group("/users")
	user.GET("", userHandler.GetUsers)
	user.GET("/:id", userHandler.GetUserById)
	user.POST("/", userHandler.CreateUser)
}
