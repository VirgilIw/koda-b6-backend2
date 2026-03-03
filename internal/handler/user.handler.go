package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/dto"
	"github.com/virgilIw/koda-b6-backend2/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) GetUsers(ctx *gin.Context) {
	data := u.service.GetUsers()

	ctx.JSON(http.StatusOK, gin.H{
		"success": "ok",
		"message": "get datas success",
		"data":    data,
	})
}

func (u *UserHandler) GetUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid id",
		})
		return
	}

	data, err := u.service.GetUserById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "get user success",
		"data":    data,
	})
}

func (u *UserHandler) CreateUser(ctx *gin.Context) {
	var req dto.CreateUserRequestDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	data := u.service.CreateUser(req)

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "user created successfully",
		"data":    data,
	})
}

func (u *UserHandler) UpdateByEmail(ctx *gin.Context) {
	queryEmail := ctx.Query("email")

	var req dto.UpdateUserRequestDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	_, err := u.service.UpdateByEmail(queryEmail, req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user updated successfully",
	})
}

func (u *UserHandler) DeleteUserById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	u.service.DeleteUserById(id)

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user deleted",
	})
}
