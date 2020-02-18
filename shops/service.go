package shops

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type service struct {
	shopStore       aumo.ShopStore
	userStore       aumo.UserStore
	shopOwnersStore aumo.ShopOwnersStore
}

//New returns an instance of `aumo.ShopService`
func New(shopStore aumo.ShopStore, shopOwnersStore aumo.ShopOwnersStore, userStore aumo.UserStore) aumo.ShopService {
	return &service{
		shopStore:       shopStore,
		shopOwnersStore: shopOwnersStore,
		userStore:       userStore,
	}
}

func (ss *service) Shop(id uint, withOwners bool) (*aumo.Shop, error) {
	return ss.shopStore.FindByID(nil, id, withOwners)
}

func (ss *service) Shops() ([]aumo.Shop, error) {
	return ss.shopStore.FindAll(nil)
}

func (ss *service) Create(s *aumo.Shop) error {
	return ss.shopStore.Save(nil, s)
}

func (ss *service) Update(id uint, s *aumo.Shop) error {
	return ss.shopStore.Update(nil, id, s)
}

func (ss *service) Delete(id uint) error {
	return ss.shopStore.Delete(nil, id)
}

func (ss *service) AddOwner(so *aumo.ShopOwners) error {
	db := ss.shopStore.DB()
	return aumo.TxDo(context.Background(), db, func(tx sqlbuilder.Tx) error {
		user, err := ss.userStore.FindByID(tx, so.UserID, false)
		if err != nil {
			return err
		}

		if user.Role == aumo.Customer {
			user.Role = aumo.ShopOwner
			err := ss.userStore.Update(tx, so.UserID, user)
			if err != nil {
				return err
			}
		}

		return ss.shopOwnersStore.Save(tx, so)
	})
}

func (ss *service) RemoveOwner(so *aumo.ShopOwners) error {
	db := ss.shopStore.DB()
	return aumo.TxDo(context.Background(), db, func(tx sqlbuilder.Tx) error {
		user, err := ss.userStore.FindByID(tx, so.UserID, true)
		if err != nil {
			return err
		}

		if user.Role == aumo.ShopOwner && len(user.Shops) == 1 {
			user.Role = aumo.Customer
			err := ss.userStore.Update(tx, so.UserID, user)
			if err != nil {
				return err
			}
		}

		return ss.shopOwnersStore.Delete(tx, so)
	})
}
