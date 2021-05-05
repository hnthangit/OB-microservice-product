package request

import "ob-product/internal/entity"

type ChargeRequest struct {
	amount     *entity.Money
	creditCard *entity.CreditCardInfo
}
