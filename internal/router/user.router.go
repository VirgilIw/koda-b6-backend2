package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/handler"
)

func RouterUser(r *gin.Engine, userHandler *handler.UserHandler) {
	users := r.Group("/users")
	users.GET("", userHandler.GetUsers)
	users.GET("/:id", userHandler.GetUserById)
	users.POST("", userHandler.CreateUser)
	users.PATCH("/email", userHandler.UpdateByEmail)
	users.DELETE("/:id", userHandler.DeleteUserById)
}
