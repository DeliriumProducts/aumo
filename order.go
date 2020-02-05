package aumo

import (
	"errors"

	"upper.io/db.v3/lib/sqlbuilder"
)

var (
	// ErrOrderProductNotFound is an error for when a user places an order on a product that doesn't exist
	ErrOrderProductNotFound = errors.New("aumo: can't place an order for a non existing product")
	// ErrOrderUserNotFound is an error for when a user doesn't exist when placing an order
	ErrOrderUserNotFound = errors.New("aumo: can't place an order for a user that doesn't exist")
	// ErrOrderNotFound when a receipt isn't found
	ErrOrderNotFound = errors.New("aumo: order not found")
)

// Order is an order in aumo
type Order struct {
	OrderID   uint     `json:"order_id" db:"order_id,omitempty"`
	UserID    uint     `json:"user_id" db:"user_id"`
	ProductID uint     `json:"product_id" db:"product_id"`
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
	Update(id uint, o *Order) error
	Delete(id uint) error
	PlaceOrder(uID, pID uint) (*Order, error)
}

// OrderStore contains all `Order`
// related persistence logic
type OrderStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id uint) (*Order, error)
	FindAll(tx Tx) ([]Order, error)
	Save(tx Tx, o *Order) error
	Update(tx Tx, id uint, o *Order) error
	Delete(tx Tx, id uint) error
}
