package shops

import (
	"github.com/deliriumproducts/aumo"
)

type service struct {
	shopStore       aumo.ShopStore
	shopOwnersStore aumo.ShopOwnersStore
}

//New returns an instance of `aumo.ShopService`
func New(shopStore aumo.ShopStore, shopOwnersStore aumo.ShopOwnersStore) aumo.ShopService {
	return &service{
		shopStore:       shopStore,
		shopOwnersStore: shopOwnersStore,
	}
}

func (ss *service) Shop(id uint) (*aumo.Shop, error) {
	return ss.shopStore.FindByID(nil, id, false)
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

func (ss *service) Owners(id uint) ([]aumo.User, error) {
	shop, err := ss.shopStore.FindByID(nil, id, true)
	if err != nil {
		return nil, err
	}

	owners := shop.Owners

	return owners, nil
}

func (ss *service) AddOwner(so *aumo.ShopOwners) error {
	return ss.shopOwnersStore.Save(nil, so)
}

func (ss *service) RemoveOwner(so *aumo.ShopOwners) error {
	return ss.shopOwnersStore.Delete(nil, so)
}
