package ordering

import (
	"context"
	"errors"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

type service struct {
	store aumo.OrderStore
	ps    aumo.ProductStore
	us    aumo.UserStore
}

// New returns an instance of `aumo.OrderService`
func New(store aumo.OrderStore, ps aumo.ProductStore, us aumo.UserStore) aumo.OrderService {
	return &service{
		store: store,
		ps:    ps,
		us:    us,
	}
}

func (o *service) Order(id string) (*aumo.Order, error) {
	return o.store.FindByID(nil, id)
}

func (o *service) Orders() ([]aumo.Order, error) {
	return o.store.FindAll(nil)
}

func (o *service) Update(id string, order *aumo.Order) error {
	return o.store.Update(nil, id, order)
}

func (o *service) Delete(id string) error {
	return o.store.Delete(nil, id)
}

func (o *service) PlaceOrder(uID string, pID uint) (*aumo.Order, error) {
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
