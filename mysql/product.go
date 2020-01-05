package mysql

import (
	"github.com/deliriumproducts/aumo"
	"github.com/jinzhu/gorm"
)

type productService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) aumo.ProductService {
	return &productService{
		db: db,
	}
}

func (p *productService) Product(id uint) *aumo.Product {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Products() ([]aumo.Product, error) {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Save(_ *aumo.Product) error {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Update(_ *aumo.Product) error {
	panic("not implemented") // TODO: Implement
}

func (p *productService) Delete(_ *aumo.Product) error {
	panic("not implemented") // TODO: Implement
}
