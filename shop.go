package aumo

import "upper.io/db.v3/lib/sqlbuilder"

// Shop is a shop in aumo
type Shop struct {
	ID     uint   `json:"shop_id" db:"shop_id,omitempty"`
	Name   string `json:"name" db:"name"`
	Owners []User `json:"owners" db:"-"`
}

// NewShop is a constructor for `Shop`
func NewShop(name string) *Shop {
	return &Shop{
		Name:   name,
		Owners: []User{},
	}
}

// ShopService contains all `Shop`
// related business logic
type ShopService interface {
	Shop(id uint) (*Shop, error)
	Shops() ([]Shop, error)
	Owners(id uint) ([]User, error)
	Update(id uint, o *Shop) error
	Delete(id uint) error
	Create(*Shop) (error)
}

// ShopStore contains all `Shop`
// related persistence logic
type ShopStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id uint, relations bool) (*Shop, error)
	FindAll(tx Tx) ([]Shop, error)
	Save(tx Tx, s *Shop) error
	Update(tx Tx, id uint, s *Shop) error
	Delete(tx Tx, id uint) error
}