package aumo

import (
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"upper.io/db.v3/lib/sqlbuilder"
)

const (
	// UserStartingPoints is the starting points of a user
	UserStartingPoints = float64(5000)
	// UserPointsPerReceipt is the amount of points a user gets per receipt claimed
	UserPointsPerReceipt = float64(500)
)

// User represents a user of aumo
type User struct {
	ID         uuid.UUID `json:"id,omitempty" db:"id,omitempty"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	Avatar     string    `json:"avatar" db:"avatar"`
	Points     float64   `json:"points" db:"points"`
	Role       Role      `json:"role" db:"role"`
	Orders     []Order   `json:"orders" db:"-"`
	Receipts   []Receipt `json:"receipts" db:"-"`
	IsVerified bool      `json:"is_verified" db:"verified"`
	Shops      []Shop    `json:"shops,omitempty" db:"-"`
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
		ID:         uuid.New(),
		Name:       name,
		Email:      strings.ToLower(strings.Trim(email, " ")),
		Password:   string(pwd),
		Avatar:     avatar,
		Points:     UserStartingPoints,
		Role:       Customer,
		Orders:     []Order{},
		Receipts:   []Receipt{},
		Shops:      []Shop{},
		IsVerified: false,
	}, nil
}

// UserService contains all `User`
// related business logic
type UserService interface {
	User(id string, relations bool) (*User, error)
	UserByEmail(email string, relations bool) (*User, error)
	Users() ([]User, error)
	Create(*User) error
	Update(id string, u *User) error
	EditRole(id string, role Role) error
	AddPoints(id string, points float64) error
	SubPoints(id string, points float64) error
	Verify(id string) error
	Delete(id string) error
}

// UserStore contains all `User`
// related persistence logic
type UserStore interface {
	DB() sqlbuilder.Database
	FindByID(tx Tx, id string, relations bool) (*User, error)
	FindByEmail(tx Tx, email string, relations bool) (*User, error)
	FindAll(tx Tx) ([]User, error)
	Save(tx Tx, u *User) error
	Update(tx Tx, id string, u *User) error
	Delete(tx Tx, id string) error
}
