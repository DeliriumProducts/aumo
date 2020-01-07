package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type userService struct {
	db sqlbuilder.Database
	rs aumo.ReceiptService
}

// NewUserService returns a mysql instance of `aumo.ProductService`
func NewUserService(db sqlbuilder.Database, rs aumo.ReceiptService) aumo.UserService {
	return &userService{
		db: db,
		rs: rs,
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

func (u *userService) ClaimReceipt(us *aumo.User, r *aumo.Receipt) error {
	r.SetUser(us.ID)
	us.ClaimReceipt(*r)
	return u.rs.Update(r.ID, r)
}
