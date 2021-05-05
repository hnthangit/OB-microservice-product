package service

import (
	"ob-product/internal/entity"
	"ob-product/internal/repository"
)

type ProductInfraService struct {
	ProductRepository repository.IProductRepository
}

func (i *ProductInfraService) SearchProduct(name string) []*entity.Product {
	return i.ProductRepository.GetProductByName(name)
}

func (i *ProductInfraService) GetProduct(id string) *entity.Product {
	return i.ProductRepository.GetProductById(id)
}

func (i *ProductInfraService) ListProducts() []*entity.Product {
	return i.ProductRepository.GetProducts()
}
