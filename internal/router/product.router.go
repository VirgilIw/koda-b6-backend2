package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/handler"
)

func RouterProduct(r *gin.Engine, productHandler *handler.ProductHandler) {
	products := r.Group("/products")

	products.GET("", productHandler.GetProducts)
	products.GET("/:id", productHandler.GetProductById)
	products.POST("", productHandler.CreateProduct)
	products.PATCH("/:id", productHandler.UpdateProduct)
	products.DELETE("/:id", productHandler.DeleteProduct)
}
