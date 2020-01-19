package mysql

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

// OrderTable is the MySQL table for holding orders
const OrderTable = "orders"

type orderStore struct {
	db sqlbuilder.Database
}

// NewOrderStore returns a mysql instance of `aumo.OrderStore`
func NewOrderStore(db sqlbuilder.Database) aumo.OrderStore {
	return &orderStore{
		db: db,
	}
}

func (o *orderStore) DB() sqlbuilder.Database {
	return o.db
}

func (o *orderStore) FindByID(tx aumo.Tx, id uint) (*aumo.Order, error) {
	var err error
	order := &aumo.Order{}

	if tx == nil {
		tx, err = o.db.NewTx(context.Background())

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

	return order, tx.Collection(OrderTable).Find("id", id).One(order)
}

func (o *orderStore) FindAll(tx aumo.Tx) ([]aumo.Order, error) {
	var err error
	orders := []aumo.Order{}

	if tx == nil {
		tx, err = o.db.NewTx(context.Background())

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

	return orders, tx.Collection(OrderTable).Find().All(&orders)
}

func (o *orderStore) Save(tx aumo.Tx, os *aumo.Order) error {
	var err error

	if tx == nil {
		tx, err = o.db.NewTx(context.Background())

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

	return tx.Collection(OrderTable).InsertReturning(os)
}

func (o *orderStore) Update(tx aumo.Tx, id uint, or *aumo.Order) error {
	var err error

	if tx == nil {
		tx, err = o.db.NewTx(context.Background())

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

	return tx.Collection(OrderTable).Find("id", id).Update(or)
}

func (o *orderStore) Delete(tx aumo.Tx, id uint) error {
	var err error

	if tx == nil {
		tx, err = o.db.NewTx(context.Background())

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

	return tx.Collection(OrderTable).Find("id", id).Delete()
}
