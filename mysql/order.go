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

func (o *orderStore) FindByID(id uint) (*aumo.Order, error) {
	order := &aumo.Order{}
	return order, o.db.Collection(OrderTable).Find("id", id).One(order)
}

func (o *orderStore) FindAll() ([]aumo.Order, error) {
	orders := []aumo.Order{}
	return orders, o.db.Collection(OrderTable).Find().All(&orders)
}

func (o *orderStore) Save(os *aumo.Order) error {
	return o.db.Collection(OrderTable).InsertReturning(os)
}

func (o *orderStore) Update(id uint, or *aumo.Order) error {
	return o.db.Collection(OrderTable).Find("id", id).Update(or)
}

func (o *orderStore) Delete(id uint) error {
	return o.db.Collection(OrderTable).Find("id", id).Delete()
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
