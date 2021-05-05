package response

import "ob-product/internal/entity"

type SearchProductResponse struct {
	results []*entity.Product
}
