package aumo

import "upper.io/db.v3/lib/sqlbuilder"

// Shop is a shop in aumo
type Shop struct {
	ID   uint   `json:"shop_id" db:"shop_id,omitempty"`
	Name string `json:"name" db:"name"`
}

// NewShop is a constructor for `Shop`
func NewShop(name string) *Shop {
	return &Shop{
		Name: name,
	}
}

// ShopService contains all `Shop`
// related business logic
type ShopService interface {
	Shop(id uint) (*Order, error)
	Owners(id uint) ([]User, error)
	Shops() ([]Order, error)
	Update(id uint, o *Order) error
	Delete(id uint) error
	New(uID, pID uint) (*Order, error)
}

// ShopStore contains all `Shop`
// related persistence logic
type ShopStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id uint) (*Order, error)
	FindAll(tx Tx) ([]Order, error)
	Save(tx Tx, o *Order) error
	Update(tx Tx, id uint, o *Order) error
	Delete(tx Tx, id uint) error
}
