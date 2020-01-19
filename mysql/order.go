package mysql

import (
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
		tx, err = o.db.NewTx(nil)

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

	return order, tx.Collection(OrderTable).Find("id", id).One(order)
}

func (o *orderStore) FindAll(tx aumo.Tx) ([]aumo.Order, error) {
	var err error
	orders := []aumo.Order{}

	if tx == nil {
		tx, err = o.db.NewTx(nil)

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

	return orders, tx.Collection(OrderTable).Find().All(&orders)
}

func (o *orderStore) Save(tx aumo.Tx, os *aumo.Order) error {
	var err error

	if tx == nil {
		tx, err = o.db.NewTx(nil)

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

	return tx.Collection(OrderTable).InsertReturning(os)
}

func (o *orderStore) Update(tx aumo.Tx, id uint, or *aumo.Order) error {
	var err error

	if tx == nil {
		tx, err = o.db.NewTx(nil)

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

	return tx.Collection(OrderTable).Find("id", id).Update(or)
}

func (o *orderStore) Delete(tx aumo.Tx, id uint) error {
	var err error

	if tx == nil {
		tx, err = o.db.NewTx(nil)

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

	return tx.Collection(OrderTable).Find("id", id).Delete()
}

// func (o *orderStore) PlaceOrder(uID, pID uint) (*aumo.Order, error) {
// 	return nil, nil
// 	// tx, err := o.db.NewTx(nil)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// product, err := u.ps.Product(pID)
// 	// if err != nil {
// 	// 	if errors.Is(err, db.ErrNoMoreRows) {
// 	// 		return nil, aumo.ErrOrderProductNotFound
// 	// 	}

// 	// 	return nil, err
// 	// }

// 	// product.DecrementStock()
// 	// o := aumo.NewOrder(user, product)
// 	// err = u.os.Create(o)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// // NOTE: is there a race condition here???
// 	// err = user.PlaceOrder(o)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// err = u.ps.Update(pID, product)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// err = u.Update(uID, user)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// return o, nil
// }
