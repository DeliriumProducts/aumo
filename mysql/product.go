package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

const ProductTable = "products"

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
	pd := &aumo.Product{}
	return pd, p.db.Collection(ProductTable).Find("id", id).One(pd)
}

func (p *productService) Products() ([]aumo.Product, error) {
	var pds []aumo.Product
	return pds, p.db.Collection(ProductTable).Find().All(&pds)
}

func (p *productService) Create(pd *aumo.Product) error {
	return p.db.Collection(ProductTable).InsertReturning(pd)
}

func (p *productService) Update(id uint, pd *aumo.Product) error {
	return p.db.Collection(ProductTable).Find("id", id).Update(pd)
}

func (p *productService) Delete(id uint) error {
	return p.db.Collection(ProductTable).Find("id", id).Delete()
}
