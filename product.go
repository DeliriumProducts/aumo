package aumo

import (
	"upper.io/db.v3/lib/sqlbuilder"
)

// Product is a product in the shop of aumo
type Product struct {
	ID          uint    `json:"id" db:"id,omitempty"`
	Name        string  `json:"name" db:"name"`
	Price       float64 `json:"price" db:"price"`
	Image       string  `json:"image" db:"image"`
	Description string  `json:"description" db:"description" `
	Stock       uint    `json:"stock" db:"stock"`
	ShopID      uint    `json:"shop_id" db:"shop_id"`
	Shop        *Shop   `json:"shop,omitempty" db:"-"`
}

// DecrementStock decreases the stock of a `Product`
func (p *Product) DecrementStock() {
	p.Stock--
}

// IncrementStock increases the stock of a `Product`
func (p *Product) IncrementStock() {
	p.Stock++
}

// NewProduct is a constructor for `Product`
func NewProduct(name string, price float64, image string, description string, stock uint, sID uint) *Product {
	return &Product{
		Name:        name,
		Price:       price,
		Image:       image,
		Description: description,
		Stock:       stock,
		ShopID:      sID,
	}
}

// ProductService contains all `Product`
// related business logic
type ProductService interface {
	Product(id uint) (*Product, error)
	Products() ([]Product, error)
	Create(*Product) error
	Update(id uint, p *Product) error
	Delete(id uint) error
	ProductsByShopID(sID uint) ([]Product, error)
}

// ProductStore contains all `Product`
// related persistence logic
type ProductStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id uint) (*Product, error)
	FindByShopID(tx Tx, shopID uint) ([]Product, error)
	FindAll(tx Tx) ([]Product, error)
	Save(tx Tx, p *Product) error
	Update(tx Tx, id uint, p *Product) error
	Delete(tx Tx, id uint) error
}
