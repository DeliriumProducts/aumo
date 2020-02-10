package mysql

import (
	"context"
	"errors"

	"github.com/deliriumproducts/aumo"
	upper "upper.io/db.v3"
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
	shop := &aumo.Shop{}

	if tx == nil {
		tx, err = p.db.NewTx(context.Background())

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	err = tx.Collection(ProductTable).Find("id", id).One(product)

	switch {
	case err == nil:
		break
	case errors.Is(err, upper.ErrNoMoreRows):
		return nil, aumo.ErrProductNotFound
	default:
		return nil, err
	}

	err = tx.Select("shops.*").
	From("shops").
	Join("products as p").On("p.shop_id = shops.shop_id").
	Where("p.shop_id = ? ", id)
	.One(shop)
	if err != nil {
		return nil, err
	}

	product.Shop = shop

	return product, nil
}

func (p *productStore) FindAll(tx aumo.Tx) ([]aumo.Product, error) {
	var err error
	products := []aumo.Product{}

	if tx == nil {
		tx, err = p.db.NewTx(context.Background())

		if err != nil {
			return nil, err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return products, tx.Collection(ProductTable).Find().All(&products)
}

func (p *productStore) Save(tx aumo.Tx, pd *aumo.Product) error {
	var err error

	if tx == nil {
		tx, err = p.db.NewTx(context.Background())

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return tx.Collection(ProductTable).InsertReturning(pd)
}

func (p *productStore) Update(tx aumo.Tx, id uint, pd *aumo.Product) error {
	var err error

	if tx == nil {
		tx, err = p.db.NewTx(context.Background())

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return tx.Collection(ProductTable).Find("id", id).Update(pd)
}

func (p *productStore) Delete(tx aumo.Tx, id uint) error {
	var err error

	if tx == nil {
		tx, err = p.db.NewTx(context.Background())

		if err != nil {
			return err
		}

		defer func() {
			if p := recover(); p != nil {
				err = tx.Rollback()
				panic(p)
			}

			if err != nil {
				err = tx.Rollback()
				return
			}

			err = tx.Commit()
		}()
	}

	return tx.Collection(ProductTable).Find("id", id).Delete()
}
