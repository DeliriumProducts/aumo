package mysql

import (
	"github.com/deliriumproducts/aumo"
	"upper.io/db.v3/lib/sqlbuilder"
)

type userService struct {
	db sqlbuilder.Database
}

// NewUserService returns a mysql instance of `aumo.ProductService`
func NewUserService(db sqlbuilder.Database) aumo.UserService {
	return &userService{
		db: db,
	}
}

func (u *userService) User(id uint) (*aumo.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *userService) Users() ([]aumo.User, error) {
	panic("not implemented") // TODO: Implement
}

func (u *userService) Create(_ *aumo.User) error {
	panic("not implemented") // TODO: Implement
}

func (u *userService) Update(id uint, ur *aumo.User) error {
	panic("not implemented") // TODO: Implement
}

func (u *userService) Delete(id uint) error {
	panic("not implemented") // TODO: Implement
}
