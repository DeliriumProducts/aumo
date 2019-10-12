package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotSufficientPoints = errors.New("aumo: user doesn't have enough points to buy this item")
	ErrNotInStock          = errors.New("aumo: shop doesn't have enough stock of the item")
)

type User struct {
	Model
	Name     string     `json:"name" gorm:"not null"`
	Email    string     `json:"email" gorm:"unique;not null"`
	Password string     `json:"-" gorm:"not null" gob:"-"`
	Avatar   string     `json:"avatar" `
	Points   float64    `json:"points" gorm:"not null"`
	Orders   []ShopItem `json:"orders" gorm:"many2many:user_shop_item;"`
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
