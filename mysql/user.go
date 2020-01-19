package mysql

import (
	"github.com/deliriumproducts/aumo"
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
		tx, err = u.db.NewTx(nil)

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
		user, err = u.userRelations(tx, "u.id = ?", id)
	} else {
		err = tx.Collection(UserTable).Find("id", id).One(user)
		user.Receipts = []aumo.Receipt{}
		user.Orders = []aumo.Order{}
	}

	return user, err
}

func (u *userStore) FindByEmail(tx aumo.Tx, email string, relations bool) (*aumo.User, error) {
	var err error
	user := &aumo.User{}

	if tx == nil {
		tx, err = u.db.NewTx(nil)

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
		user, err = u.userRelations(tx, "u.email = ?", email)
	} else {
		err = tx.Collection(UserTable).Find("email", email).One(user)
		user.Receipts = []aumo.Receipt{}
		user.Orders = []aumo.Order{}
	}

	return user, err
}

func (u *userStore) userRelations(tx aumo.Tx, where string, args ...interface{}) (*aumo.User, error) {
	var err error

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
		userReceipts = []UserReceipt{}
		orders       = []OrderProduct{}
	)

	err = u.db.
		Select("u.id", "u.name", "u.email", "u.password", "u.avatar", "u.points", "u.role", "r.receipt_id", "r.content", "r.user_id").
		From("users as u").
		Join("receipts as r").On("u.id = r.user_id").
		Where(where, args).
		All(&userReceipts)
	if err != nil {
		return nil, err
	}

	err = u.db.
		Select("o.user_id", "o.product_id", "p.name", "p.description", "p.price", "p.image", "p.price", "p.image", "p.id", "p.stock", "o.order_id").
		From("users as u").
		Join("orders as o").On("u.id = o.user_id").
		Join("products as p").On("o.product_id = p.id").
		Where(where, args).
		All(&orders)
	if err != nil {
		return nil, err
	}

	user := &userReceipts[0].User
	user.Orders = []aumo.Order{}
	user.Receipts = []aumo.Receipt{}

	for _, o := range orders {
		ord := o.Order
		ord.Product = &o.Product

		user.Orders = append(user.Orders, ord)
	}
	for _, r := range userReceipts {
		user.Receipts = append(user.Receipts, r.Receipt)
	}

	return user, nil
}

func (u *userStore) FindAll(tx aumo.Tx) ([]aumo.User, error) {
	var err error
	users := []aumo.User{}

	if tx == nil {
		tx, err = u.db.NewTx(nil)

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
		tx, err = u.db.NewTx(nil)

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

	return tx.Collection(UserTable).InsertReturning(us)
}

func (u *userStore) Update(tx aumo.Tx, id uint, ur *aumo.User) error {
	var err error

	if tx == nil {
		tx, err = u.db.NewTx(nil)

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
		tx, err = u.db.NewTx(nil)

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
