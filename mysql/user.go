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
	user := &aumo.User{}

	var err error

	if relations {
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
			Select("u.id", "u.name", "u.email", "u.password", "u.avatar", "u.points", "r.receipt_id", "r.content", "r.user_id").
			From("users as u").
			Join("receipts as r").On("u.id = r.user_id").
			Where("u.id = ?", id).
			All(&userReceipts)
		if err != nil {
			return nil, err
		}

		err = u.db.
			Select("o.user_id", "o.product_id", "p.name", "p.description", "p.price", "p.image", "p.price", "p.image", "p.id", "p.stock", "o.order_id").
			From("users as u").
			Join("orders as o").On("u.id = o.user_id").
			Join("products as p").On("o.product_id = p.id").
			Where("u.id = ?", id).
			All(&orders)
		if err != nil {
			return nil, err
		}

		user = &userReceipts[0].User
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
	} else {
		err = u.db.Collection(UserTable).Find("id", id).One(user)
		user.Receipts = []aumo.Receipt{}
		user.Orders = []aumo.Order{}
	}

	return user, err
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

func (u *userService) ClaimReceipt(user *aumo.User, rID uint) error {
	receipt, err := u.rs.Receipt(rID)
	if err != nil {
		return err
	}

	// NOTE: is there a race condition here???
	err = receipt.Claim(user.ID)
	if err != nil {
		return err
	}

	user.ClaimReceipt(receipt)
	return u.rs.Update(rID, receipt)
}

func (u *userService) PlaceOrder(user *aumo.User, pID uint) error {
	product, err := u.ps.Product(pID)
	if err != nil {
		return err
	}

	product.DecrementStock()
	o := aumo.NewOrder(user, product)
	err = u.os.Create(o)
	if err != nil {
		return err
	}

	// NOTE: is there a race condition here???
	err = user.PlaceOrder(o)
	if err != nil {
		return err
	}

	err = u.ps.Update(pID, product)
	if err != nil {
		return err
	}

	err = u.Update(user.ID, user)
	if err != nil {
		return err
	}

	return nil
}
