package mysql

import (
	"github.com/deliriumproducts/aumo"
	"golang.org/x/net/context"
	"upper.io/db.v3/lib/sqlbuilder"
)

// ShopTable is the MySQL table for holding orders
const ShopTable = "shops"

type shopStore struct {
	db sqlbuilder.Database
}

func NewShopStore(db sqlbuilder.Database) aumo.ShopStore {
	return &shopStore{
		db: db,
	}
}

func (s *shopStore) DB() sqlbuilder.Database {
	return s.db
}

func (s *shopStore) FindByID(tx aumo.Tx, id uint) (*aumo.Shop, error) {
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

	return shop, tx.Collection(ShopTable).Find("id", id).One(shop)
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

	return tx.Collection(OrderTable).InsertReturning(ss)
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

	return tx.Collection(OrderTable).Find("id", id).Update(sp)
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

	return tx.Collection(OrderTable).Find("id", id).Delete()
}

//func (s *shopStore) shopRelations(tx aumo.Tx, where string, args ...interface{}) (*aumo.User, error) {
//	var err error
//	shop := &aumo.Shop{}
//
//	type ShopUsers struct {
//		aumo.User `db:",inline"`
//	}
//
//	var (
//		users = []ShopUsers{}
//	)
//
//	err = tx.
//		Select("*").
//		From(ProductTable).
//		Where(where, args).
//		One(shop)
//	if err != nil {
//		return nil, err
//	}
//
//	err = tx.Select("u.user_id", "u.name", "u.email", "u.avatar", "u.points").
//		From(ShopTable).
//		Join("users as u").On("shops_owners.user_id = u.user_id").
//		Where(where, args).
//		All(&users)
//	if err != nil {
//		return nil, err
//	}
//
//}
