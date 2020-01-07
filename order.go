package aumo

type Order struct {
	ID        uint    `json:"id" db:"id,omitempty"`
	UserID    uint    `json:"userID" db:"user_id"`
	ProductID uint    `json:"productID" db:"product_id"`
	Product   Product `json:"product" db:"-"`
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
