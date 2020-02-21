package mysql

import (
	"context"
	"errors"

	"github.com/deliriumproducts/aumo"
	"github.com/go-sql-driver/mysql"
	upper "upper.io/db.v3"
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

func (u *userStore) FindByID(tx aumo.Tx, id string, relations bool) (*aumo.User, error) {
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
		user.Shops = []aumo.Shop{}
	}

	switch {
	case err == nil:
		break
	case errors.Is(err, upper.ErrNoMoreRows):
		return nil, aumo.ErrUserNotFound
	default:
		return nil, err
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
		user.Shops = []aumo.Shop{}
	}

	switch {
	case err == nil:
		break
	case errors.Is(err, upper.ErrNoMoreRows):
		return nil, aumo.ErrUserNotFound
	default:
		return nil, err
	}

	return user, err
}

func (u *userStore) userRelations(tx aumo.Tx, where string, args ...interface{}) (*aumo.User, error) {
	var err error
	user := &aumo.User{}

	type (
		OrderProduct struct {
			aumo.Order   `db:",inline"`
			aumo.Product `db:",inline"`
		}

		ReceiptShop struct {
			aumo.Receipt `db:",inline"`
			aumo.Shop    `db:",inline"`
		}
	)
	var (
		orders   = []OrderProduct{}
		receipts = []ReceiptShop{}
		shops    = []aumo.Shop{}
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
		Select("*").
		From(UserTable).
		Join("receipts as r").On("users.id = r.user_id").
		Join("shops as s").On("r.shop_id = s.shop_id").
		Where(where, args).
		All(&receipts)
	if err != nil {
		return nil, err
	}

	err = tx.
		Select("o.*", "p.*").
		From(UserTable).
		Join("orders as o").On("users.id = o.user_id").
		Join("products as p").On("o.product_id = p.id").
		Where(where, args).
		All(&orders)
	if err != nil {
		return nil, err
	}

	// Get shops only if the user is an admin or a shop owner
	if user.Role == aumo.ShopOwner {
		err = tx.Select("s.*").
			From("shop_owners").
			Join("shops as s").On("shop_owners.shop_id = s.shop_id").
			Join("users").On("users.id = shop_owners.user_id").
			Where(where, args).
			All(&shops)
		if err != nil {
			return nil, err
		}
	}

	if user.Role == aumo.Admin {
		err = tx.
			Collection(ShopTable).
			Find().
			All(&shops)
		if err != nil {
			return nil, err
		}
	}

	user.Orders = []aumo.Order{}
	user.Receipts = []aumo.Receipt{}
	user.Shops = shops

	for i := range receipts {
		receipt := receipts[i].Receipt
		receipt.Shop = &receipts[i].Shop
		receipt.ShopID = receipts[i].Shop.ID
		receipt.Shop.Owners = []aumo.User{}
		user.Receipts = append(user.Receipts, receipt)
	}

	for i := range orders {
		order := orders[i].Order
		order.Product = &orders[i].Product
		user.Orders = append(user.Orders, order)
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

	_, err = tx.Collection(UserTable).Insert(us)
	if mysqlError, ok := err.(*mysql.MySQLError); ok {
		if mysqlError.Number == ErrDupEntry {
			return aumo.ErrDuplicateEmail
		}
	}

	return err
}

func (u *userStore) Update(tx aumo.Tx, id string, ur *aumo.User) error {
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

func (u *userStore) Delete(tx aumo.Tx, id string) error {
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
