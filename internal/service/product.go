package service

import (
	"ob-product/internal/entity"
)

type IProductService interface {
	SearchProduct(name string) []*entity.Product
	GetProduct(id string) *entity.Product
	ListProducts() []*entity.Product
}
