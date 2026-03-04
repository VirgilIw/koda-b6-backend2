package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	data := p.service.GetProducts()

	ctx.JSON(http.StatusOK, gin.H{
		"success": "ok",
		"message": "get data product success",
		"data":    data,
	})
}
