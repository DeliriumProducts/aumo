package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type productService struct {
	db sqlbuilder.Database
}

// NewProductService returns a mysql instance of `aumo.ProductService`
func NewProductService(db sqlbuilder.Database) aumo.ProductService {
	return &productService{
		db: db,
	}
}

func (p *productService) Product(id uint) (*aumo.Product, error) {
	var pd *aumo.Product
	return pd, p.db.Collection("products").Find("id", id).One(pd)
}

func (p *productService) Products() ([]aumo.Product, error) {
	var pds []aumo.Product
	return pds, p.db.Collection("products").Find().All(&pds)
}

func (p *productService) Create(pd *aumo.Product) error {
	return p.db.Collection("products").InsertReturning(pd)
}

func (p *productService) Update(id uint, pd *aumo.Product) error {
	return p.db.Collection("products").Find("id", id).Update(pd)
}

func (p *productService) Delete(id uint, pd *aumo.Product) error {
	return p.db.Collection("products").Find("id", id).Delete()
}
