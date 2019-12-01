package aumo

import (
	"github.com/fr3fou/aumo/server/aumo/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser creates a user
func (a *Aumo) CreateUser(name, email, password, avatar string) (models.User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return models.User{}, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(pwd),
		Avatar:   avatar,
		Points:   5000,
		Orders:   []models.ShopItem{},
		Receipts: []models.Receipt{},
	}

	if err := a.db.Create(user).Error; err != nil {
		return models.User{}, err
	}

	return *user, nil
}

// GetUserByEmail returns a user that has a matching email
func (a *Aumo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := a.firstX(&user, "email = ?", email)
	return user, err
}

// GetUserByID returns a user that has a matching id
func (a *Aumo) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := a.firstX(&user, "id = ?", id)
	return user, err
}

// UpdateUser updates a user
func (a *Aumo) UpdateUser(u models.User) (models.User, error) {
	return u, a.updateX(&u)
}

// DeleteUser deletes a user
func (a *Aumo) DeleteUser(i models.User) error {
	return a.deleteX(i)
}

// BuyUserShopItem calls BuytItem on the user struct, decrements
// the stock of the shop item then it updates it
func (a *Aumo) BuyUserShopItem(u models.User, si models.ShopItem, quantity uint) error {
	err := u.BuyItem(si, quantity)
	if err != nil {
		return err
	}

	si.DecrementStock(quantity)

	err = a.updateX(u)
	if err != nil {
		return err
	}

	err = a.updateX(si)
	if err != nil {
		return err
	}

	return nil
}
