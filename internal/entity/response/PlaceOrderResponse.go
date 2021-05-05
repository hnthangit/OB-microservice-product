package response

import "ob-product/internal/entity"

type PlaceOrderResponse struct {
	order *entity.OrderResult
}
