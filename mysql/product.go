package mysql

import (
	"database/sql"

	"github.com/deliriumproducts/aumo"
)

type productService struct {
	db *sql.DB
}

// NewProductService returns a mysql instance of `aumo.ProductService`
func NewProductService(db *sql.DB) aumo.ProductService {
	return &productService{
		db: db,
	}
}

func (p *productService) Product(id uint) (*aumo.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Products() ([]aumo.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Create(_ *aumo.Product) error {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Update(_ *aumo.Product) error {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Delete(_ *aumo.Product) error {
	panic("not implemented") // TODO: Implement
}
