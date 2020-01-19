package ordering

import "github.com/deliriumproducts/aumo"

type orderService struct {
	store aumo.OrderStore
}

// New returns an instance of `aumo.OrderService`
func New(store aumo.OrderStore) aumo.OrderService {
	return &orderService{
		store: store,
	}
}
