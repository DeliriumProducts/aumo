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

func (p *productStore) DB() sqlbuilder.Database {
	return p.db
}

func (p *productStore) FindByID(tx aumo.Tx, id uint) (*aumo.Product, error) {
	var err error
	product := &aumo.Product{}

	if tx == nil {
		tx, err = p.db.NewTx(nil)

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return product, tx.Collection(ProductTable).Find("id", id).One(product)
}

func (p *productStore) FindAll(tx aumo.Tx) ([]aumo.Product, error) {
	var err error
	products := []aumo.Product{}

	if tx == nil {
		tx, err = p.db.NewTx(nil)

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return products, tx.Collection(ProductTable).Find().All(&products)
}

func (p *productStore) Save(tx aumo.Tx, pd *aumo.Product) error {
	var err error

	if tx == nil {
		tx, err = p.db.NewTx(nil)

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return tx.Collection(ProductTable).InsertReturning(pd)
}

func (p *productStore) Update(tx aumo.Tx, id uint, pd *aumo.Product) error {
	var err error

	if tx == nil {
		tx, err = p.db.NewTx(nil)

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return tx.Collection(ProductTable).Find("id", id).Update(pd)
}

func (p *productStore) Delete(tx aumo.Tx, id uint) error {
	var err error

	if tx == nil {
		tx, err = p.db.NewTx(nil)

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				tx.Rollback()
				panic(p)
			} else if err != nil {
				tx.Rollback()
			} else {
				err = tx.Commit()
			}
		}()
	}

	return tx.Collection(ProductTable).Find("id", id).Delete()
}
