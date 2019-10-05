package aumo

import (
	"errors"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotSufficientPoints = errors.New("aumo: user doesn't have enough points to buy this item")
	ErrNotInStock          = errors.New("aumo: shop doesn't have enough stock of the item")
)

type User struct {
	gorm.Model
	Name     string     `json:"name" gorm:"not null"`
	Email    string     `json:"email" gorm:"unique;not null"`
	Password string     `json:"-" gorm:"not null"`
	Points   float64    `json:"points" gorm:"not null"`
	Orders   []ShopItem `gorm:"many2many:user_shop_item;"`
	Receipts []Receipt  `json:"receipts"`
}

// ValidatePassword checks if the passed password is the correct one
func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	return false
}

// BuyItem adds the passed ShopItem to the user's inventory
// if they have enough money to buy the desired quantity;
// substracts points from the user
func (u *User) BuyItem(si ShopItem, quantity uint) error {
	pointsToPay := si.Price * float64(quantity)
	// Check if the user has enough points
	if u.Points-pointsToPay < 0 {
		return ErrNotSufficientPoints
	}

	// Check if there is enough in stock
	if si.Stock-quantity < 0 {
		return ErrNotInStock
	}

	// Substract the points of the user
	u.Points -= pointsToPay

	// Add the item to the orders array
	u.Orders = append(u.Orders, si)
	return nil
}

// ClaimReceipt claims a receipt and adds it to the receipts array
func (u *User) ClaimReceipt(r Receipt) {
	u.Receipts = append(u.Receipts, r)
}

// CreateUser creates a user
func (a *Aumo) CreateUser(name, email, password string) (User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return User{}, err
	}

	user := &User{
		Name:     name,
		Email:    email,
		Password: string(pwd),
		Points:   5000,
		Orders:   []ShopItem{},
		Receipts: []Receipt{},
	}

	if err := a.db.Create(user).Error; err != nil {
		return User{}, err
	}

	return *user, nil
}

// GetUserByEmail returns a user that has a matching email
func (a *Aumo) GetUserByEmail(email string) (User, error) {
	var user User
	err := a.firstX(&user, "email = ?", email)
	return user, err
}

// GetUserByID returns a user that has a matching id
func (a *Aumo) GetUserByID(id uint) (User, error) {
	var user User
	err := a.firstX(&user, "id = ?", id)
	return user, err
}

// UpdateUser updates a user
func (a *Aumo) UpdateUser(u User) (User, error) {
	return u, a.updateX(&u)
}

// DeleteUser deletes a user
func (a *Aumo) DeleteUser(i User) error {
	return a.deleteX(i)
}

// BuyUserShopItem calls BuytItem on the user struct, decrements
// the stock of the shop item then it updates it
func (a *Aumo) BuyUserShopItem(u User, si ShopItem, quantity uint) error {
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
