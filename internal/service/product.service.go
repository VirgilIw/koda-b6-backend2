package service

import (
	"github.com/virgilIw/koda-b6-backend2/internal/dto"
	"github.com/virgilIw/koda-b6-backend2/internal/model"
	"github.com/virgilIw/koda-b6-backend2/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repository *repository.ProductRepository) *ProductService {
	return &ProductService{
		repo: repository,
	}
}

// Get all products
func (p *ProductService) GetProducts() []dto.ProductsDto {
	data := p.repo.GetProducts()

	var result []dto.ProductsDto

	for _, v := range data {
		result = append(result, dto.ProductsDto{
			Id:          v.Id,
			ProductName: v.ProductName,
			Rating:      v.Rating,
		})
	}

	return result
}

// Get product by ID
func (p *ProductService) GetProductById(id int) dto.ProductsDto {
	data, found := p.repo.GetProductById(id)

	if !found {
		return dto.ProductsDto{}
	}

	return dto.ProductsDto{
		Id:          data.Id,
		ProductName: data.ProductName,
		Rating:      data.Rating,
	}
}

// Create product
func (p *ProductService) CreateProduct(req dto.ProductsDto) dto.ProductsDto {

	product := model.ProductModel{
		ProductName: req.ProductName,
		Rating:      req.Rating,
	}

	data := p.repo.CreateProduct(product)

	return dto.ProductsDto{
		Id:          data.Id,
		ProductName: data.ProductName,
		Rating:      data.Rating,
	}
}

// Update product
func (p *ProductService) UpdateProduct(id int, req dto.ProductsDto) dto.ProductsDto {

	product := model.ProductModel{
		ProductName: req.ProductName,
		Rating:      req.Rating,
	}

	data, found := p.repo.UpdateProduct(id, product)

	if !found {
		return dto.ProductsDto{}
	}

	return dto.ProductsDto{
		Id:          data.Id,
		ProductName: data.ProductName,
		Rating:      data.Rating,
	}
}

// Delete product
func (p *ProductService) DeleteProduct(id int) dto.ProductsDto {
	data := p.repo.DeleteProduct(id)

	if data.Id == 0 {
		return dto.ProductsDto{}
	}

	return dto.ProductsDto{
		Id:          data.Id,
		ProductName: data.ProductName,
		Rating:      data.Rating,
	}
}
