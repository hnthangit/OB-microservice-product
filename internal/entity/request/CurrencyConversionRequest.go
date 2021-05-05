package request

import "ob-product/internal/entity"

type CurrencyConversionRequest struct {
	from   *entity.Money
	toCode string
}
