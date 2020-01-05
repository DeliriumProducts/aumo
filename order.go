package aumo

// Order is an order of a `User`
type Order struct {
	User    *User    `json:"-"`
	Product *Product `json:"product"`
	Amount  uint     `json:"amount"`
}

// NewOrder is a constructor for `Order`
func NewOrder(u *User, p *Product, n uint) *Order {
	return &Order{
		User:    u,
		Product: p,
		Amount:  n,
	}
}
