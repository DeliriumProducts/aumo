package aumo

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// UserStartingPoints is the starting points of a user
const UserStartingPoints = float64(5000)

var (
	// ErrNotSufficientPoints is an error for when the user doens't have enough points
	ErrNotSufficientPoints = errors.New("aumo: user doesn't have enough points to buy this item")
	// ErrNotInStock is an error for when an item isn't in stock
	ErrNotInStock = errors.New("aumo: shop doesn't have enough stock of the item")
)

// User represents a user of aumo
type User struct {
	ID       uint      `json:"-" db:"id,omitempty"`
	Name     string    `json:"name" db:"name"`
	Email    string    `json:"email" db:"email"`
	Password string    `json:"-" db:"password"`
	Avatar   string    `json:"avatar" db:"avatar"`
	Points   float64   `json:"points" db:"points"`
	Orders   []Order   `json:"orders" db:"-"`
	Receipts []Receipt `json:"receipts" db:"-"`
}

// ClaimReceipt claims a receipt and adds it to the receipts array
func (u *User) ClaimReceipt(r *Receipt) {
	u.Receipts = append(u.Receipts, *r)
}

// ValidatePassword checks if the passed password is the correct one
func (u *User) ValidatePassword(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		return true
	}

	return false
}

// PlaceOrder adds the passed Product to the user's inventory
// if they have enough money to buy the desired quantity;
// subtracts points from the user
func (u *User) PlaceOrder(o *Order) error {
	p := o.Product

	// Check if the user has enough points
	if u.Points-p.Price < 0 {
		return ErrNotSufficientPoints
	}

	// Check if there is enough in stock
	if p.Stock < 1 {
		return ErrNotInStock
	}

	// Subtract the points of the user
	u.Points -= p.Price

	// Add the item to the orders array
	u.Orders = append(u.Orders, *o)
	return nil
}

// NewUser is a constructor for `User`
func NewUser(name string, email string, password string, avatar string) (*User, error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), 12)

	if err != nil {
		return nil, err
	}

	return &User{
		Name:     name,
		Email:    email,
		Password: string(pwd),
		Avatar:   avatar,
		Points:   UserStartingPoints,
		Orders:   []Order{},
		Receipts: []Receipt{},
	}, nil
}

// UserService contains all `User`
// related business logic
type UserService interface {
	User(id uint, relations bool) (*User, error)
	UserByEmail(email string, relations bool) (*User, error)
	Users() ([]User, error)
	Create(*User) error
	Update(id uint, u *User) error
	Delete(id uint) error
	ClaimReceipt(u *User, rID uint) (*Receipt, error)
	PlaceOrder(u *User, pid uint) error
}
