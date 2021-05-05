package repository

import "ob-product/internal/entity"

type IProductRepository interface {
	GetProductById(id string) *entity.Product
	GetProducts() []*entity.Product
	GetProductByName(name string) []*entity.Product
}
