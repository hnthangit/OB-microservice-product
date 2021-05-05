package repository

import (
	"ob-product/internal/entity"
	"ob-product/internal/infra/file"
	"strings"
)

type ProductInfraRepository struct{}

func (r *ProductInfraRepository) GetProductById(id string) *entity.Product {
	var found *entity.Product

	var products, err = file.ReadCatalogFile()

	if err == nil {
		for i := 0; i < len(products); i++ {
			if id == products[i].Id {
				found = products[i]
			}
		}
	}

	return found
}

func (r *ProductInfraRepository) GetProducts() []*entity.Product {
	var products, err = file.ReadCatalogFile()
	if err == nil && products != nil {
		return products
	}

	return products
}

func (r *ProductInfraRepository) GetProductByName(name string) []*entity.Product {

	var found []*entity.Product

	var products, err = file.ReadCatalogFile()
	if err == nil {
		for i := 0; i < len(products); i++ {
			if strings.Contains(strings.ToLower(products[i].Name), strings.ToLower(name)) ||
				strings.Contains(strings.ToLower(products[i].Description), strings.ToLower(name)) {
				found = append(found, products[i])
			}
		}
	}

	return found
}
