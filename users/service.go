package users

import "github.com/deliriumproducts/aumo"

type service struct {
	store aumo.UserStore
}

// New returns an instance of `aumo.UserService`
func New(store aumo.UserStore) aumo.UserService {
	return &service{
		store: store,
	}
}

func (us *service) User(id uint, relations bool) (*aumo.User, error) {
	return us.store.FindByID(id, relations)
}

func (us *service) UserByEmail(email string, relations bool) (*aumo.User, error) {
	return us.store.FindByEmail(email, relations)
}

func (us *service) Users() ([]aumo.User, error) {
	return us.store.FindAll()
}

func (us *service) Create(u *aumo.User) error {
	return us.store.Save(u)
}

func (us *service) Update(id uint, u *aumo.User) error {
	return us.store.Update(id, u)
}

func (us *service) Delete(id uint) error {
	return us.store.Delete(id)
}
