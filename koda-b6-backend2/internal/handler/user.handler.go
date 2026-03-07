package handler

import (
	"fmt"
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

func (h *UserHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	users, err := h.service.GetUsers(ctx)
	fmt.Println(users)
	fmt.Println("ERROR:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response[string]{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response[[]dto.UserDataResponse]{
		Success: true,
		Message: "Success get users",
		Result:  users,
	})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response[string]{
			Success: false,
			Message: "invalid id",
		})
		return
	}

	data, err := h.service.GetUserById(ctx, id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response[string]{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	if id != data.Id {
		c.JSON(http.StatusNotFound, dto.Response[string]{
			Success: false,
			Message: "data not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response[dto.UserDataResponse]{
		Success: true,
		Message: "Success get user",
		Result:  data,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var data dto.CreateUserRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response[string]{
			Success: false,
			Message: "internal server error",
		})
		return
	}

	result, err := h.service.CreateUser(ctx, data.FullName, data.Email, data.Password)

	if err != nil {

	}
	c.JSON(http.StatusCreated, dto.Response[dto.CreateUserResponse]{
		Success: true,
		Message: "created user success",
		Result:  result,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response[string]{
			Success: false,
			Message: "invalid user id",
		})
		return
	}

	err = h.service.DeleteUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response[string]{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response[string]{
		Success: true,
		Message: "user deleted successfully",
	})
}
