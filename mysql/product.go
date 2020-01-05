package mysql

import (
	"github.com/deliriumproducts/aumo"
)

type productService struct {
	db *DB
}

// NewProductService returns a mysql instance of `aumo.ProductService`
func NewProductService(db *DB) aumo.ProductService {
	return &productService{
		db: db,
	}
}

func (p *productService) Product(id uint) (*aumo.Product, error) {
	pd := &aumo.Product{}
	return pd, p.db.First(pd, id).Error
}

func (p *productService) Products() ([]aumo.Product, error) {
	pds := []aumo.Product{}
	return pds, p.db.Find(&pds).Error
}

func (p *productService) Create(pd *aumo.Product) error {
	return p.db.Create(pd).Error
}

func (p *productService) Update(pd *aumo.Product) error {
	return p.db.Model(pd).Updates(pd).Error
}

func (p *productService) Delete(pd *aumo.Product) error {
	return p.db.Unscoped().Delete(pd).Error
}
