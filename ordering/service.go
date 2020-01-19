package ordering

import (
	"context"
	"errors"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type orderService struct {
	store aumo.OrderStore
	ps    aumo.ProductStore
	us    aumo.UserStore
}

// New returns an instance of `aumo.OrderService`
func New(store aumo.OrderStore, ps aumo.ProductStore, us aumo.UserStore) aumo.OrderService {
	return &orderService{
		store: store,
		ps:    ps,
		us:    us,
	}
}

func (o *orderService) Order(id uint) (*aumo.Order, error) {
	return o.store.FindByID(nil, id)
}

func (o *orderService) Orders() ([]aumo.Order, error) {
	return o.store.FindAll(nil)
}

func (o *orderService) Update(id uint, order *aumo.Order) error {
	return o.store.Update(nil, id, order)
}

func (o *orderService) Delete(id uint) error {
	return o.store.Delete(nil, id)
}

func (o *orderService) PlaceOrder(uID, pID uint) (*aumo.Order, error) {
	order := &aumo.Order{}
	err := aumo.TxDo(context.Background(), o.store.DB(), func(tx sqlbuilder.Tx) error {
		// Get Product
		product, err := o.ps.FindByID(tx, pID)
		if err != nil {
			if errors.Is(err, db.ErrNoMoreRows) {
				return aumo.ErrOrderProductNotFound
			}

			return err
		}

		// Get User
		user, err := o.us.FindByID(tx, uID, false)
		if err != nil {
			if errors.Is(err, db.ErrNoMoreRows) {
				return aumo.ErrOrderUserNotFound
			}

			return err
		}

		// Create order
		order = aumo.NewOrder(user, product)

		// Place order
		err = user.PlaceOrder(order)
		if err != nil {
			return err
		}

		// Decrement stock
		product.DecrementStock()

		// Save order
		err = o.store.Save(tx, order)
		if err != nil {
			return err
		}

		// Update product
		err = o.ps.Update(tx, pID, product)
		if err != nil {
			return err
		}

		// Update user
		err = o.us.Update(tx, uID, user)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return order, nil
}
