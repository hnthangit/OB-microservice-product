package request

import "ob-product/internal/entity"

type SendOrderConfimationRequest struct {
	email string
	order *entity.OrderResult
}
