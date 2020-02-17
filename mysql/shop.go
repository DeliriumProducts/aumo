package mysql

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

// ShopTable is the MySQL table for holding shops
const ShopTable = "shops"

// ShopOwnersTable is the MySQL table for holding relations
// between users and shops
const ShopOwnersTable = "shop_owners"

type shopStore struct {
	db sqlbuilder.Database
}

type shopOwnersStore struct {
	db sqlbuilder.Database
}

// NewShopStore returns a mysql instance of `aumo.ShopStore`
func NewShopStore(db sqlbuilder.Database) aumo.ShopStore {
	return &shopStore{
		db: db,
	}
}

// NewShopOwnersStore returns a mysql instance of `aumo.ShopOwnersStore`
func NewShopOwnersStore(db sqlbuilder.Database) aumo.ShopOwnersStore {
	return &shopOwnersStore{
		db: db,
	}
}

func (s *shopStore) DB() sqlbuilder.Database {
	return s.db
}

func (s *shopStore) FindByID(tx aumo.Tx, id uint, relations bool) (*aumo.Shop, error) {
	var err error
	shop := &aumo.Shop{}

	if tx == nil {
		tx, err = s.db.NewTx(context.Background())

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

	if relations {
		shop, err = s.shopRelations(tx, "shops.shop_id = ?", id)
	} else {
		err = tx.Collection(ShopTable).Find("id", id).One(shop)
		shop.Owners = []aumo.User{}
	}

	return shop, err
}

func (s *shopStore) FindAll(tx aumo.Tx) ([]aumo.Shop, error) {
	var err error
	shops := []aumo.Shop{}

	if tx == nil {
		tx, err = s.db.NewTx(context.Background())

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

	return shops, tx.Collection(ShopTable).Find().All(&shops)
}

func (s *shopStore) Save(tx aumo.Tx, ss *aumo.Shop) error {
	var err error

	if tx == nil {
		tx, err = s.db.NewTx(context.Background())

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

	return tx.Collection(ShopTable).InsertReturning(ss)
}

func (s *shopStore) Update(tx aumo.Tx, id uint, sp *aumo.Shop) error {
	var err error

	if tx == nil {
		tx, err = s.db.NewTx(context.Background())

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

	return tx.Collection(ShopTable).Find("id", id).Update(sp)
}

func (s *shopStore) Delete(tx aumo.Tx, id uint) error {
	var err error

	if tx == nil {
		tx, err = s.db.NewTx(context.Background())

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

	return tx.Collection(ShopTable).Find("id", id).Delete()
}

func (s *shopStore) shopRelations(tx aumo.Tx, where string, args ...interface{}) (*aumo.Shop, error) {
	var err error
	shop := &aumo.Shop{}
	owners := []aumo.User{}

	err = tx.
		Select("*").
		From(ShopTable).
		Where(where, args).
		One(shop)
	if err != nil {
		return nil, err
	}

	err = tx.Select("u.*").
		From("shop_owners").
		Join("users as u").On("shops_owners.user_id = u.user_id").
		Where(where, args).
		All(&owners)
	if err != nil {
		return nil, err
	}

	shop.Owners = owners

	return shop, nil
}

func (s *shopOwnersStore) Save(tx aumo.Tx, so *aumo.ShopOwners) error {
	var err error

	if tx == nil {
		tx, err = s.db.NewTx(context.Background())

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

	return tx.Collection(ShopOwnersTable).InsertReturning(so)
}

func (s *shopOwnersStore) Delete(tx aumo.Tx, so *aumo.ShopOwners) error {
	var err error

	if tx == nil {
		tx, err = s.db.NewTx(context.Background())

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

	_, err = tx.DeleteFrom(ShopOwnersTable).
		Where("shop_owners.shop_id = ? AND shop_owners.user_id = ?", so.ShopID, so.UserID).
		Exec()

	return err
}
