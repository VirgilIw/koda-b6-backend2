package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/dto"
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

// Get all products
func (p *ProductHandler) GetProducts(ctx *gin.Context) {
	data := p.service.GetProducts()

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "get data product success",
		"data":    data,
	})
}

// Get product by id
func (p *ProductHandler) GetProductById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	data := p.service.GetProductById(id)

	if data.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    data,
	})
}

// Create product
func (p *ProductHandler) CreateProduct(ctx *gin.Context) {
	var req dto.ProductsDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := p.service.CreateProduct(req)

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "product created",
		"data":    data,
	})
}

// Update product
func (p *ProductHandler) UpdateProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	var req dto.ProductsDto

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	data := p.service.UpdateProduct(id, req)

	if data.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "product updated",
		"data":    data,
	})
}

// Delete product
func (p *ProductHandler) DeleteProduct(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	data := p.service.DeleteProduct(id)

	if data.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "product not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "product deleted",
		"data":    data,
	})
}
