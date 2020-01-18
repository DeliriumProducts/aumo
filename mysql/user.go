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

// NewUserStore returns a mysql instance of `aumo.UserStore`
func NewUserStore(db sqlbuilder.Database) aumo.UserStore {
	return &userStore{
		db: db,
	}
}

func (u *userStore) FindByID(id uint, relations bool) (*aumo.User, error) {
	user := &aumo.User{}
	var err error

	if relations {
		user, err = u.userRelations("u.id = ?", id)
	} else {
		err = u.db.Collection(UserTable).Find("id", id).One(user)
		user.Receipts = []aumo.Receipt{}
		user.Orders = []aumo.Order{}
	}

	return user, err
}

func (u *userStore) FindByEmail(email string, relations bool) (*aumo.User, error) {
	user := &aumo.User{}
	var err error

	if relations {
		user, err = u.userRelations("u.email = ?", email)
	} else {
		err = u.db.Collection(UserTable).Find("email", email).One(user)
		user.Receipts = []aumo.Receipt{}
		user.Orders = []aumo.Order{}
	}

	return user, err
}

func (u *userStore) userRelations(where string, args ...interface{}) (*aumo.User, error) {
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

func (u *userStore) FindAll() ([]aumo.User, error) {
	uss := []aumo.User{}
	return uss, u.db.Collection(UserTable).Find().All(&uss)
}

func (u *userStore) Save(us *aumo.User) error {
	return u.db.Collection(UserTable).InsertReturning(us)
}

func (u *userStore) Update(id uint, ur *aumo.User) error {
	return u.db.Collection(UserTable).Find("id", id).Update(ur)
}

func (u *userStore) Delete(id uint) error {
	return u.db.Collection(UserTable).Find("id", id).Delete()
}
