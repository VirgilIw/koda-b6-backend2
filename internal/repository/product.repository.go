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
	db *[]model.ProductModel
}

func NewProductRepository(db *[]model.ProductModel) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

// Get all products
func (p *ProductRepository) GetProducts() []model.ProductModel {
	return *p.db
}

// Get product by ID
func (p *ProductRepository) GetProductById(id int) (model.ProductModel, bool) {
	for _, product := range p.db {
		if product.Id == id {
			return product, true
		}
	}
	return model.ProductModel{}, false
}

// Create product
func (p *ProductRepository) CreateProduct(product model.ProductModel) model.ProductModel {
	product.Id = len(*p.db) + 1
	*p.db = append(*p.db, product)

	return product
}

// Update product
func (p *ProductRepository) UpdateProduct(id int, updated model.ProductModel) (model.ProductModel, bool) {
	for i, product := range p.db {
		if product.Id == id {
			p.db[i].ProductName = updated.ProductName
			p.db[i].Rating = updated.Rating
			return p.db[i], true
		}
	}
	return model.ProductModel{}, false
}

// Delete product
func (p *ProductRepository) DeleteProduct(id int) model.ProductModel {
	var deleted model.ProductModel
	newData := []model.ProductModel{}

	for _, product := range *p.db {
		if product.Id == id {
			deleted = product
			continue
		}
		newData = append(newData, product)
	}

	*p.db = newData
	return deleted
}
