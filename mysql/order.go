package mysql

import (
	"context"
	"errors"

	"github.com/deliriumproducts/aumo"
	upper "upper.io/db.v3"
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

func (o *orderStore) FindByID(tx aumo.Tx, id string) (*aumo.Order, error) {
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

	err = tx.Collection(OrderTable).Find("id", id).One(order)

	switch {
	case err == nil:
		break
	case errors.Is(err, upper.ErrNoMoreRows):
		return nil, aumo.ErrOrderNotFound
	default:
		return nil, err
	}

	return order, err
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

	_, err = tx.Collection(OrderTable).Insert(os)
	return err
}

func (o *orderStore) Update(tx aumo.Tx, id string, or *aumo.Order) error {
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

func (o *orderStore) Delete(tx aumo.Tx, id string) error {
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
