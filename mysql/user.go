package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/deliriumproducts/aumo"
	"github.com/go-sql-driver/mysql"
	"upper.io/db.v3/lib/sqlbuilder"
)

// UserTable is the MySQL table for holding users
const UserTable = "users"

type userStore struct {
	db sqlbuilder.Database
}

func (u *userStore) DB() sqlbuilder.Database {
	return u.db
}

// NewUserStore returns a mysql instance of `aumo.UserStore`
func NewUserStore(db sqlbuilder.Database) aumo.UserStore {
	return &userStore{
		db: db,
	}
}

func (u *userStore) FindByID(tx aumo.Tx, id uint, relations bool) (*aumo.User, error) {
	var err error
	user := &aumo.User{}

	if tx == nil {
		tx, err = u.db.NewTx(context.Background())

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
		user, err = u.userRelations(tx, "users.id = ?", id)
	} else {
		err = tx.Collection(UserTable).Find("id", id).One(user)
		user.Receipts = []aumo.Receipt{}
		user.Orders = []aumo.Order{}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, aumo.ErrUserNotFound
	}

	return user, err
}

func (u *userStore) FindByEmail(tx aumo.Tx, email string, relations bool) (*aumo.User, error) {
	var err error
	user := &aumo.User{}

	if tx == nil {
		tx, err = u.db.NewTx(context.Background())

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
		user, err = u.userRelations(tx, "users.email = ?", email)
	} else {
		err = tx.Collection(UserTable).Find("email", email).One(user)
		user.Receipts = []aumo.Receipt{}
		user.Orders = []aumo.Order{}
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, aumo.ErrUserNotFound
	}

	return user, err
}

func (u *userStore) userRelations(tx aumo.Tx, where string, args ...interface{}) (*aumo.User, error) {
	var err error
	user := &aumo.User{}

	type (
		UserReceipt struct {
			aumo.User    `db:",inline"`
			aumo.Receipt `db:",inline"`
		}
		OrderProduct struct {
			aumo.Order   `db:",inline"`
			aumo.Product `db:",inline"`
		}
	)
	var (
		receipts = []UserReceipt{}
		orders   = []OrderProduct{}
	)

	err = tx.
		Select("*").
		From(UserTable).
		Where(where, args).
		One(user)
	if err != nil {
		return nil, err
	}

	err = tx.
		Select("r.receipt_id", "r.content", "r.user_id").
		From(UserTable).
		Join("receipts as r").On("users.id = r.user_id").
		Where(where, args).
		All(&receipts)
	if err != nil {
		return nil, err
	}

	err = tx.
		Select("o.user_id", "o.product_id", "p.name", "p.description", "p.price", "p.image", "p.price", "p.image", "p.id", "p.stock", "o.order_id").
		From(UserTable).
		Join("orders as o").On("users.id = o.user_id").
		Join("products as p").On("o.product_id = p.id").
		Where(where, args).
		All(&orders)
	if err != nil {
		return nil, err
	}

	user.Orders = []aumo.Order{}
	user.Receipts = []aumo.Receipt{}

	for i := range orders {
		order := orders[i].Order
		order.Product = &orders[i].Product
		user.Orders = append(user.Orders, order)
	}

	for i := range receipts {
		user.Receipts = append(user.Receipts, receipts[i].Receipt)
	}

	return user, nil
}

func (u *userStore) FindAll(tx aumo.Tx) ([]aumo.User, error) {
	var err error
	users := []aumo.User{}

	if tx == nil {
		tx, err = u.db.NewTx(context.Background())

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

	return users, tx.Collection(UserTable).Find().All(&users)
}

func (u *userStore) Save(tx aumo.Tx, us *aumo.User) error {
	var err error

	if tx == nil {
		tx, err = u.db.NewTx(context.Background())

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

	err = tx.Collection(UserTable).InsertReturning(us)
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == ErrDupEntry {
			return aumo.ErrDuplicateEmail
		}
	}

	return err
}

func (u *userStore) Update(tx aumo.Tx, id uint, ur *aumo.User) error {
	var err error

	if tx == nil {
		tx, err = u.db.NewTx(context.Background())

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

	return tx.Collection(UserTable).Find("id", id).Update(ur)
}

func (u *userStore) Delete(tx aumo.Tx, id uint) error {
	var err error

	if tx == nil {
		tx, err = u.db.NewTx(context.Background())

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

	return tx.Collection(UserTable).Find("id", id).Delete()
}
