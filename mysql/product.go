package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

// ProductTable is the MySQL table for holding products
const ProductTable = "products"

type productStore struct {
	db sqlbuilder.Database
}

// NewProductStore returns a mysql instance of `aumo.ProductStore`
func NewProductStore(db sqlbuilder.Database) aumo.ProductStore {
	return &productStore{
		db: db,
	}
}

func (p *productStore) FindByID(id uint) (*aumo.Product, error) {
	product := &aumo.Product{}
	return product, p.db.Collection(ProductTable).Find("id", id).One(product)
}

func (p *productStore) FindAll() ([]aumo.Product, error) {
	products := []aumo.Product{}
	return products, p.db.Collection(ProductTable).Find().All(&products)
}

func (p *productStore) Save(pd *aumo.Product) error {
	return p.db.Collection(ProductTable).InsertReturning(pd)
}

func (p *productStore) Update(id uint, pd *aumo.Product) error {
	return p.db.Collection(ProductTable).Find("id", id).Update(pd)
}

func (p *productStore) Delete(id uint) error {
	return p.db.Collection(ProductTable).Find("id", id).Delete()
}
