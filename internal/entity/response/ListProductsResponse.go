package response

import "ob-product/internal/entity"

type ListProductResponse struct {
	products []*entity.Product
}
