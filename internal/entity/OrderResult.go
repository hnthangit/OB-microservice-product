package entity

type OrderResult struct {
	orderId            string
	shippingTrackingId string
	shippingCost       *Money
	items              []*OrderItem
}
