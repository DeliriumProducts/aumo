package aumo

import "upper.io/db.v3/lib/sqlbuilder"

// Shop is a shop in aumo
type Shop struct {
	ID       uint      `json:"id" db:"shop_id,omitempty"`
	Name     string    `json:"name" db:"name"`
	Image    string    `json:"image" db:"image"`
	Owners   []User    `json:"owners,omitempty" db:"-"`
	Products []Product `json:"products,omitempty" db:"-"`
}

// NewShop is a constructor for `Shop`
func NewShop(name, image string) *Shop {
	return &Shop{
		Name:   name,
		Image:  image,
		Owners: []User{},
	}
}

// ShopService contains all `Shop`
// related business logic
type ShopService interface {
	Shop(id uint, withOwners bool) (*Shop, error)
	Shops() ([]Shop, error)
	AddOwner(id uint, email string) error
	RemoveOwner(id uint, email string) error
	Update(id uint, o *Shop) error
	Delete(id uint) error
	Create(*Shop) error
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

// ShopOwnersStore currently contains some "shop_owners table"
// related persistence logic
type ShopOwnersStore interface {
	Save(tx Tx, sID uint, uID string) error
	Delete(tx Tx, sID uint, uID string) error
	DeleteByUser(tx Tx, uID string) error
}
