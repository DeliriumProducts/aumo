package mysql

import (
	"context"
	"errors"

	"github.com/deliriumproducts/aumo"
	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3"
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

func (s *shopStore) FindByID(tx aumo.Tx, id uint, withOwners bool) (*aumo.Shop, error) {
	var err error
	shop := &aumo.Shop{}
	products := []aumo.Product{}
	owners := []aumo.User{}

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

	err = tx.Collection(ShopTable).Find("shop_id", id).One(shop)
	if errors.Is(err, db.ErrNoMoreRows) {
		return nil, aumo.ErrShopNotFound
	}

	if withOwners {
		err = tx.Select("u.*").
			From("shop_owners").
			Join("users as u").On("shop_owners.user_id = u.id").
			Where("shop_owners.shop_id", id).
			All(&owners)

		if err != nil {
			return nil, err
		}
	}

	err = tx.
		Collection(ProductTable).
		Find("shop_id", id).
		All(&products)
	if err != nil {
		return nil, err
	}

	shop.Owners = owners
	shop.Products = products

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

	return tx.Collection(ShopTable).Find("shop_id", id).Update(sp)
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

	return tx.Collection(ShopTable).Find("shop_id", id).Delete()
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

	err = tx.Collection(ShopOwnersTable).InsertReturning(so)
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == ErrBadRef {
			return aumo.ErrUserNotFound
		}
		if mysqlError.Number == ErrDupEntry {
			return aumo.ErrUserAlreadyOwnsShop
		}
	}

	return err
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
