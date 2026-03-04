package router

import (
	"github.com/gin-gonic/gin"
	"github.com/virgilIw/koda-b6-backend2/internal/handler"
)

func RouterProduct(r *gin.Engine, productHandler *handler.ProductHandler) {
	products := r.Group("/products")
	products.GET("", productHandler.GetProducts)
}
