package mysql

import (
	"context"

	"github.com/deliriumproducts/aumo"
	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
)

type shopOwnersStore struct {
	db sqlbuilder.Database
}

// NewShopOwnersStore returns a mysql instance of `aumo.ShopOwnersStore`
func NewShopOwnersStore(db sqlbuilder.Database) aumo.ShopOwnersStore {
	return &shopOwnersStore{
		db: db,
	}
}

func (s *shopOwnersStore) Save(tx aumo.Tx, sID uint, uID string) error {
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

	_, err = tx.
		InsertInto(ShopOwnersTable).
		Columns(
			"shop_id", "user_id",
		).
		Values(
			sID, uID,
		).
		Exec()
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

func (s *shopOwnersStore) Delete(tx aumo.Tx, sID uint, uID string) error {
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
		Where("shop_owners.shop_id = ? AND shop_owners.user_id = ?", sID, uID).
		Exec()

	return err
}

func (s *shopOwnersStore) DeleteByUser(tx aumo.Tx, uID string) error {
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
		Where("shop_owners.user_id", uID).
		Exec()

	return err
}
