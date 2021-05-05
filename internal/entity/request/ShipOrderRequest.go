package request

import "ob-product/internal/entity"

type ShipOrderRequest struct {
	address *entity.Address
	items   []*entity.CartItem
}
