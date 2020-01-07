package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type userService struct {
	db sqlbuilder.Database
	rs aumo.ReceiptService
	ps aumo.ProductService
	os aumo.OrderService
}

// NewUserService returns a mysql instance of `aumo.ProductService`
func NewUserService(db sqlbuilder.Database, rs aumo.ReceiptService, ps aumo.ProductService, os aumo.OrderService) aumo.UserService {
	return &userService{
		db: db,
		rs: rs,
		ps: ps,
		os: os,
	}
}

func (u *userService) User(id uint) (*aumo.User, error) {
	us := &aumo.User{}
	return us, u.db.Collection("users").Find("id", id).One(us)
}

func (u *userService) Users() ([]aumo.User, error) {
	uss := []aumo.User{}
	return uss, u.db.Collection("users").Find().All(&uss)
}

func (u *userService) Create(us *aumo.User) error {
	return u.db.Collection("users").InsertReturning(us)
}

func (u *userService) Update(id uint, ur *aumo.User) error {
	return u.db.Collection("users").Find("id", id).Update(ur)
}

func (u *userService) Delete(id uint) error {
	return u.db.Collection("users").Find("id", id).Delete()
}

func (u *userService) ClaimReceipt(us *aumo.User, rid uint) error {
	r, err := u.rs.Receipt(rid)
	if err != nil {
		return err
	}

	// NOTE: is there a race condition here???
	err = r.SetUser(us.ID)
	if err != nil {
		return err
	}

	us.ClaimReceipt(*r)
	return u.rs.Update(r.ID, r)
}

func (u *userService) PlaceOrder(us *aumo.User, pid uint) error {
	p, err := u.ps.Product(pid)
	if err != nil {
		return err
	}

	o := aumo.NewOrder(us.ID, pid, p)
	// NOTE: is there a race condition here???
	err = us.PlaceOrder(*o)
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
