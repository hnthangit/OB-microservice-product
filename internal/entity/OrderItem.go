package entity

type OrderItem struct {
	item *CartItem
	cost *Money
}
