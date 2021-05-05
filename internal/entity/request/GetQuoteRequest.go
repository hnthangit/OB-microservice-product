package request

import "ob-product/internal/entity"

type GetQuoteRequest struct {
	address *entity.Address
	items   []*entity.CartItem
}
