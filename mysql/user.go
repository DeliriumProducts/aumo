package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

const UserTable = "users"

type userService struct {
	db sqlbuilder.Database
	rs aumo.ReceiptService
	ps aumo.ProductService
	os aumo.OrderService
}

// NewUserService returns a mysql instance of `aumo.UserService`
func NewUserService(db sqlbuilder.Database, rs aumo.ReceiptService, ps aumo.ProductService, os aumo.OrderService) aumo.UserService {
	return &userService{
		db: db,
		rs: rs,
		ps: ps,
		os: os,
	}
}

func (u *userService) User(id uint, relations bool) (*aumo.User, error) {
	us := &aumo.User{}

	var err error

	if relations {
		type UserReceipt struct {
			aumo.User    `db:",inline"`
			aumo.Receipt `db:",inline"`
		}

		urs := []UserReceipt{}
		orders := []aumo.Order{}
		err = u.db.
			Select("u.id", "u.name", "u.email", "u.password", "u.avatar", "u.points", "r.receipt_id", "r.content", "r.user_id").
			From("users as u").
			Join("receipts as r").On("u.id = r.user_id").
			Where("u.id = ?", id).
			All(&urs)
		if err != nil {
			return nil, err
		}

		err = u.db.
			Select("o.user_id", "o.product_id", "p.name", "p.description", "p.price", "p.image", "p.price", "p.image", "p.id", "o.order_id").
			From("users as u").
			Join("orders as o").On("u.id = o.user_id").
			Join("products as p").On("o.product_id = p.id").
			Where("u.id = ?", id).
			All(&orders)
		if err != nil {
			return nil, err
		}

		us = &urs[0].User
		us.Orders = orders
		for _, r := range urs {
			us.Receipts = append(us.Receipts, r.Receipt)
		}
	} else {
		err = u.db.Collection(UserTable).Find("id", id).One(us)
		us.Receipts = []aumo.Receipt{}
		us.Orders = []aumo.Order{}
	}

	return us, err
}

func (u *userService) Users() ([]aumo.User, error) {
	uss := []aumo.User{}
	return uss, u.db.Collection(UserTable).Find().All(&uss)
}

func (u *userService) Create(us *aumo.User) error {
	return u.db.Collection(UserTable).InsertReturning(us)
}

func (u *userService) Update(id uint, ur *aumo.User) error {
	return u.db.Collection(UserTable).Find("id", id).Update(ur)
}

func (u *userService) Delete(id uint) error {
	return u.db.Collection(UserTable).Find("id", id).Delete()
}

func (u *userService) ClaimReceipt(us *aumo.User, rid uint) error {
	r, err := u.rs.Receipt(rid)
	if err != nil {
		return err
	}

	// NOTE: is there a race condition here???
	err = r.Claim(us.ID)
	if err != nil {
		return err
	}

	us.ClaimReceipt(r)
	return u.rs.Update(r.ReceiptID, r)
}

func (u *userService) PlaceOrder(us *aumo.User, pid uint) error {
	p, err := u.ps.Product(pid)
	if err != nil {
		return err
	}

	o := aumo.NewOrder(us, p)
	// NOTE: is there a race condition here???
	err = us.PlaceOrder(o)
	if err != nil {
		return err
	}

	p.DecrementStock()

	err = u.os.Create(o)
	if err != nil {
		return err
	}

	err = u.ps.Update(pid, p)
	if err != nil {
		return err
	}

	return nil
}
