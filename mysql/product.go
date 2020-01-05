package mysql

import (
	"github.com/deliriumproducts/aumo"
)

type productService struct {
	db *DB
}

func NewProductService(db *DB) aumo.ProductService {
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
