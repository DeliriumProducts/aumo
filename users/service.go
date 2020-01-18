package users

import "github.com/deliriumproducts/aumo"

type userService struct {
	store aumo.UserStore
}

// New returns an instance of `aumo.UserService`
func New(store aumo.UserStore) aumo.UserService {
	return &userService{
		store: store,
	}
}

func (us *userService) User(id uint, relations bool) (*aumo.User, error) {
	return us.store.FindByID(id, relations)
}

func (us *userService) UserByEmail(email string, relations bool) (*aumo.User, error) {
	return us.store.FindByEmail(email, relations)
}

func (us *userService) Users() ([]aumo.User, error) {
	return us.store.FindAll()
}

func (us *userService) Create(u *aumo.User) error {
	return us.store.Save(u)
}

func (us *userService) Update(id uint, u *aumo.User) error {
	return us.store.Update(id, u)
}

func (us *userService) Delete(id uint) error {
	return us.store.Delete(id)
}
