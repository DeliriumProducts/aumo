package users

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type service struct {
	store aumo.UserStore
	so    aumo.ShopOwnersStore
}

// New returns an instance of `aumo.UserService`
func New(store aumo.UserStore, so aumo.ShopOwnersStore) aumo.UserService {
	return &service{
		store: store,
		so:    so,
	}
}

func (us *service) User(id string, relations bool) (*aumo.User, error) {
	return us.store.FindByID(nil, id, relations)
}

func (us *service) UserByEmail(email string, relations bool) (*aumo.User, error) {
	return us.store.FindByEmail(nil, email, relations)
}

func (us *service) Users() ([]aumo.User, error) {
	return us.store.FindAll(nil)
}

func (us *service) Create(u *aumo.User) error {
	return us.store.Save(nil, u)
}

func (us *service) Update(id string, u *aumo.User) error {
	return us.store.Update(nil, id, u)
}

func (us *service) EditRole(id string, role aumo.Role) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.Role = role

		return us.store.Update(tx, id, user)
	})
}

func (us *service) Verify(id string) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.IsVerified = true

		return us.store.Update(tx, id, user)
	})
}

func (us *service) AddPoints(id string, points float64) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.Points += points

		return us.store.Update(tx, id, user)
	})
}

func (us *service) SubPoints(id string, points float64) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.Points -= points

		return us.store.Update(tx, id, user)
	})
}

func (us *service) Delete(id string) error {
	return us.store.Delete(nil, id)
}
