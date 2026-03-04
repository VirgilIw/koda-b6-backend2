package service

import (
	"github.com/virgilIw/koda-b6-backend2/internal/dto"
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
