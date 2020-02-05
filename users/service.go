package users

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type service struct {
	store aumo.UserStore
}

// New returns an instance of `aumo.UserService`
func New(store aumo.UserStore) aumo.UserService {
	return &service{
		store: store,
	}
}

func (us *service) User(id uint, relations bool) (*aumo.User, error) {
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

func (us *service) Update(id uint, u *aumo.User) error {
	return us.store.Update(nil, id, u)
}

func (us *service) EditRole(id uint, role aumo.Role) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.Role = role

		return us.store.Update(tx, id, user)
	})
}

func (us *service) Verify(id uint) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.IsVerified = true

		return us.store.Update(tx, id, user)
	})
}

func (us *service) AddPoints(id uint, points float64) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.Points += points

		return us.store.Update(tx, id, user)
	})
}

func (us *service) SubPoints(id uint, points float64) error {
	return aumo.TxDo(context.Background(), us.store.DB(), func(tx sqlbuilder.Tx) error {
		user, err := us.store.FindByID(tx, id, false)
		if err != nil {
			return err
		}

		user.Points -= points

		return us.store.Update(tx, id, user)
	})
}

func (us *service) Delete(id uint) error {
	return us.store.Delete(nil, id)
}
