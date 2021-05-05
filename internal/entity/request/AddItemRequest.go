package request

import "ob-product/internal/entity"

type AddItemRequest struct {
	userId string
	item   *entity.CartItem
}
