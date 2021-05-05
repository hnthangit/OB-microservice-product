package response

import "ob-product/internal/entity"

type GetQuoteResponse struct {
	costUsd *entity.Money
}
