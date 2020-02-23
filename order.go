package aumo

import (
	"github.com/google/uuid"
	"upper.io/db.v3/lib/sqlbuilder"
)

// Order is an order in aumo
type Order struct {
	OrderID   uuid.UUID `json:"order_id" db:"order_id,omitempty"`
	UserID    uuid.UUID `json:"user_id" db:"user_id"`
	ProductID uint      `json:"product_id" db:"product_id"`
	Product   *Product  `json:"product" db:"-"`
}

// NewOrder is a constructor for `Order`
func NewOrder(u *User, p *Product) *Order {
	return &Order{
		OrderID:   uuid.New(),
		UserID:    u.ID,
		ProductID: p.ID,
		Product:   p,
	}
}

// OrderService contains all `Order`
// related business logic
type OrderService interface {
	Order(id string) (*Order, error)
	Orders() ([]Order, error)
	Update(id string, o *Order) error
	Delete(id string) error
	PlaceOrder(uID string, pID uint) (*Order, error)
}

// OrderStore contains all `Order`
// related persistence logic
type OrderStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id string) (*Order, error)
	FindAll(tx Tx) ([]Order, error)
	Save(tx Tx, o *Order) error
	Update(tx Tx, id string, o *Order) error
	Delete(tx Tx, id string) error
}
