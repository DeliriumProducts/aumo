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
		Points:   0,
		Orders:   []ShopItem{},
	}

	if err := a.DB.Create(user).Error; err != nil {
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
func (a *Aumo) UpdateUser(old, new User) (User, error) {
	return old, a.updateX(&old, new)
}

// DeleteUser deletes a user
func (a *Aumo) DeleteUser(i User) error {
	return a.deleteX(i)
}

// ValidatePassword checks if the passed password is the correct one
func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	return false
}

// BuyItem adds the passed ShopItem to the user's inventory
// if they have enough money to buy the desired quantity
func (u *User) BuyItem(si ShopItem, quantity uint) error {
	// Check if the user has enough points
	if u.Points-si.Price*float64(quantity) < 0 {
		return ErrNotSufficientPoints
	}

	// Check if there is enough in stock
	if si.Stock-quantity < 0 {
		return ErrNotInStock
	}

	// Substract the points of the user
	u.Points -= si.Price * float64(quantity)

	// Add the item to the orders array
	u.Orders = append(u.Orders, si)
	return nil
}

// // BuyItem takes in a shopItem and purchases it if there are enough points
// func (u *User) BuyItem(si ShopItem, quantity uint) error {
// 	if u.Points-si.Price*float64(quantity) < 0 {
// 		return ErrNotSufficientPoints
// 	}

// 	if si.Quantity-quantity < 0 {
// 		return ErrNotInStock
// 	}

// 	u.Points = u.Points - si.Price*float64(quantity)
// 	err := u.Update(*u)
// 	if err != nil {
// 		return err
// 	}

// 	si.Quantity = si.Quantity - quantity
// 	err = si.Update(si)
// 	if err != nil {
// 		return err
// 	}

// 	return u.a.DB.Model(u).Association("Orders").Append(si).Error
// }
