package repository

import "github.com/virgilIw/koda-b6-backend2/internal/model"

var productData = []model.ProductModel{
	{
		Id:          1,
		ProductName: "Kopi tubruk",
		Rating:      4.8,
	},
}

type ProductRepository struct {
	db []model.ProductModel
}

func NewProductRepository(db *[]model.ProductModel) *ProductRepository {
	return &ProductRepository{
		db: productData,
	}
}

func (p *ProductRepository) GetProducts() []model.ProductModel {
	return p.db
}
