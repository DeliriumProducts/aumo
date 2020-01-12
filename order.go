package aumo

// Order is an order in aumo
type Order struct {
	OrderID   uint     `json:"orderID" db:"order_id,omitempty"`
	UserID    uint     `json:"userID" db:"user_id"`
	ProductID uint     `json:"productID" db:"product_id"`
	Product   *Product `json:"product" db:"-"`
}

// NewOrder is a constructor for `Order`
func NewOrder(u *User, p *Product) *Order {
	return &Order{
		UserID:    u.ID,
		ProductID: p.ID,
		Product:   p,
	}
}

// OrderService contains all `Order`
// related business logic
type OrderService interface {
	Order(id uint) (*Order, error)
	Orders() ([]Order, error)
	Create(*Order) error
	Update(id uint, o *Order) error
	Delete(id uint) error
}
