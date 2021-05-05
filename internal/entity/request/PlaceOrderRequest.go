package request

import "ob-product/internal/entity"

type PlaceOrderRequest struct {
	userId       string
	userCurrency string
	address      *entity.Address
	email        string
	creditCard   *entity.CreditCardInfo
}
